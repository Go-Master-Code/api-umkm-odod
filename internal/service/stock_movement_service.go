package service

import (
	"context"
	"errors"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/payloads"
	"umkm-odod/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// interface
type StockMovementService interface {
	AddStock(ctx context.Context, req dto.AddStockRequest) (dto.StockMovementResponse, error)
	ReduceStock(ctx context.Context, req dto.ReduceStockRequest) (dto.StockMovementResponse, error)
	GetCurrentStock(ctx context.Context, itemVariantID string) (dto.CurrentStockResponse, error)
}

// struct implementasi
type stockMovementService struct {
	db   *gorm.DB // pakai db karena akan dipakai untuk transaction (1 runtutan dengan sale)
	repo repository.StockMovementRepository
}

// constructor
func NewStockMovementService(db *gorm.DB, repo repository.StockMovementRepository) StockMovementService {
	return &stockMovementService{
		db:   db,
		repo: repo,
	}
}

// struct method
func (s *stockMovementService) AddStock(ctx context.Context, req dto.AddStockRequest) (dto.StockMovementResponse, error) {
	// begin transaction
	tx := s.db.Begin()

	if tx.Error != nil {
		tx.Rollback()
		return dto.StockMovementResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	// ambil userID dari context
	userID := ctx.Value(constants.ContextUserID).(string)

	/*
		Penjelasan penting ⭐⭐⭐
		Kenapa pakai payload dulu?

		Karena nanti:

		sales service
		purchasing service
		adjustment service

		semua bisa reuse payload yang sama.
	*/

	// ========================================
	// CREATE PAYLOAD
	// ========================================

	payload := payloads.CreateStockMovementPayload{
		TenantID:      tenantID,
		ItemVariantID: req.ItemVariantID,
		MovementType:  constants.MovementPurchase, // barang masuk dari pembelian
		Qty:           req.Qty,
		ReferenceType: "",
		ReferenceID:   "",
		Notes:         req.Notes,
		CreatedBy:     userID,
	}

	// buat model untuk param repo movement
	movement := model.StockMovement{
		ID:            uuid.NewString(),
		TenantID:      payload.TenantID,
		ItemVariantID: payload.ItemVariantID,
		MovementType:  payload.MovementType,
		Qty:           payload.Qty,
		ReferenceType: payload.ReferenceType,
		ReferenceID:   payload.ReferenceID,
		Notes:         payload.Notes,
		CreatedBy:     payload.CreatedBy,
	}

	err := s.repo.CreateMovement(ctx, tx, &movement)
	if err != nil {
		tx.Rollback() // jika gagal create data, rollback
		return dto.StockMovementResponse{}, err
	}

	// commit transaction, rollback jika gagal
	// ========================================
	// COMMIT TRANSACTION
	// ========================================

	err = tx.Commit().Error

	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// get by id untuk preload data tenant dan item variant
	newMovement, err := s.repo.GetMovementByID(ctx, movement.ID)
	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// convert model to dto
	movementDTO := helper.ConvertToDTOStockMovementSingle(newMovement)
	return movementDTO, nil
}

func (s *stockMovementService) ReduceStock(ctx context.Context, req dto.ReduceStockRequest) (dto.StockMovementResponse, error) {
	// begin transaction
	tx := s.db.Begin()

	if tx.Error != nil {
		tx.Rollback()
		return dto.StockMovementResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	currentStock, err := s.repo.GetCurrentStock(ctx, tx, req.ItemVariantID)
	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// validasi stok tidak boleh negatif
	if currentStock < req.Qty {
		return dto.StockMovementResponse{}, errors.New("insufficient stock")
	}

	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	// ambil userID dari context
	userID := ctx.Value(constants.ContextUserID).(string)

	payload := payloads.CreateStockMovementPayload{
		TenantID:      tenantID,
		ItemVariantID: req.ItemVariantID,
		MovementType:  constants.MovementSale, // sale = penjualan otomatis qty jadi minus karena stok berkurang
		Qty:           -req.Qty,               // IMPORTANT ⭐⭐⭐
		ReferenceType: "",
		ReferenceID:   "",
		Notes:         req.Notes,
		CreatedBy:     userID,
	}

	movement := model.StockMovement{
		ID:            uuid.NewString(),
		TenantID:      payload.TenantID,
		ItemVariantID: payload.ItemVariantID,
		MovementType:  payload.MovementType,
		Qty:           payload.Qty,
		ReferenceType: payload.ReferenceType,
		ReferenceID:   payload.ReferenceID,
		Notes:         payload.Notes,
		CreatedBy:     payload.CreatedBy,
	}

	// eksekusi repo
	err = s.repo.CreateMovement(ctx, tx, &movement)

	if err != nil {
		tx.Rollback()
		return dto.StockMovementResponse{}, err
	}

	// ========================================
	// COMMIT TRANSACTION
	// ========================================

	err = tx.Commit().Error

	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// get by id untuk preload data tenant dan item variant
	newMovement, err := s.repo.GetMovementByID(ctx, movement.ID)
	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// convert model to dto
	movementDTO := helper.ConvertToDTOStockMovementSingle(newMovement)
	return movementDTO, nil
}

func (s *stockMovementService) GetCurrentStock(ctx context.Context, itemVariantID string) (dto.CurrentStockResponse, error) {
	// begin transaction
	tx := s.db.Begin()

	if tx.Error != nil {
		tx.Rollback()
		return dto.CurrentStockResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	stock, err := s.repo.GetCurrentStock(ctx, tx, itemVariantID)
	if err != nil {
		return dto.CurrentStockResponse{}, err
	}

	// convert to dto CurrentStockResponse
	currentStockDTO := helper.ConvertToDTOCurrentStock(itemVariantID, stock)
	return currentStockDTO, nil
}
