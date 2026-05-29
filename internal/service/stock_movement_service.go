package service

import (
	"context"
	"errors"
	"fmt"
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
	CreateAdjustment(ctx context.Context, req dto.CreateStockAdjustmentRequest) (dto.StockMovementResponse, error)
	GetStockCard(ctx context.Context, itemVariantID string) ([]dto.StockCardResponse, error)
}

// struct implementasi
type stockMovementService struct {
	db                *gorm.DB // pakai db karena akan dipakai untuk transaction (1 runtutan dengan sale)
	stockMovementRepo repository.StockMovementRepository
	itemVariantRepo   repository.ItemVariantRepository
}

// constructor
func NewStockMovementService(db *gorm.DB, stockMovementRepo repository.StockMovementRepository, itemVariantRepo repository.ItemVariantRepository) StockMovementService {
	return &stockMovementService{
		db:                db,
		stockMovementRepo: stockMovementRepo,
		itemVariantRepo:   itemVariantRepo,
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

	err := s.stockMovementRepo.CreateMovement(ctx, tx, &movement)
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
	newMovement, err := s.stockMovementRepo.GetMovementByID(ctx, tenantID, movement.ID)
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

	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	currentStock, err := s.stockMovementRepo.GetCurrentStock(ctx, tenantID, tx, req.ItemVariantID)
	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// validasi stok tidak boleh negatif
	if currentStock < req.Qty {
		return dto.StockMovementResponse{}, errors.New("insufficient stock")
	}

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
	err = s.stockMovementRepo.CreateMovement(ctx, tx, &movement)

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
	newMovement, err := s.stockMovementRepo.GetMovementByID(ctx, tenantID, movement.ID)
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

	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	stock, err := s.stockMovementRepo.GetCurrentStock(ctx, tenantID, tx, itemVariantID)
	if err != nil {
		return dto.CurrentStockResponse{}, err
	}

	// convert to dto CurrentStockResponse
	currentStockDTO := helper.ConvertToDTOCurrentStock(itemVariantID, stock)
	return currentStockDTO, nil
}

func (s *stockMovementService) CreateAdjustment(ctx context.Context, req dto.CreateStockAdjustmentRequest) (dto.StockMovementResponse, error) {
	// begin transaction
	tx := s.db.Begin()

	// cek apakah tx error
	if tx.Error != nil {
		return dto.StockMovementResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()

		if r != nil {
			tx.Rollback()
		}
	}()

	// ambil tenant ID dan user ID dari JWT
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	userID := ctx.Value(constants.ContextUserID).(string)

	// get item variant
	itemVariant, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, req.ItemVariantID)
	if err != nil {
		tx.Rollback() // wajib rollback kalau ada error pada transaction
		return dto.StockMovementResponse{}, err
	}

	// cek stok dari itemVariant tersebut
	currentStock, err := s.stockMovementRepo.GetCurrentStock(ctx, tenantID, tx, itemVariant.ID)
	if err != nil {
		tx.Rollback() // rollback lagi jika ada yang salah dengan transaction
		return dto.StockMovementResponse{}, err
	}

	// ========================================
	// TENTUKAN ARAH MOVEMENT
	// ADD = positif
	// REDUCE = negatif
	// ========================================

	// validasi adjustment qty (tidak boleh > dari stok yang tersedia)
	adjustmentQty := req.Qty // defaultnya positif (dianggap ADD)

	// validasi jika tipe adjustment adalah REDUCE
	if req.Type == constants.AdjustmentReduce {
		if currentStock < req.Qty { // jika stok yang tersedia < qty yang diminta
			tx.Rollback() // wajib rollback tiap ada case yanga tidak sesuai
			return dto.StockMovementResponse{}, errors.New("insufficient stock")
		}
		// jika stok cukup, maka ubah menjadi - (negatif) karena ini sifatnya reduce
		adjustmentQty = -req.Qty
	}

	// create stock movement
	movement := model.StockMovement{
		ID:            uuid.NewString(),
		TenantID:      tenantID,
		ItemVariantID: itemVariant.ID,
		MovementType:  req.Type,
		Qty:           adjustmentQty,
		ReferenceType: "ADJUSTMENT",
		ReferenceID:   "",
		Notes:         fmt.Sprintf("Reason: %s | Notes: %s", req.Reason, req.Notes),
		CreatedBy:     userID,
	}

	// save data movement
	err = s.stockMovementRepo.CreateMovement(ctx, tx, &movement)

	if err != nil {
		tx.Rollback() // rollback tiap kali ada kasus error
		return dto.StockMovementResponse{}, err
	}

	// commit transaction
	err = tx.Commit().Error
	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// ========================================
	// AMBIL DATA FINAL DENGAN PRELOAD
	// ========================================
	newMovement, err := s.stockMovementRepo.GetMovementByID(ctx, tenantID, movement.ID)
	if err != nil {
		return dto.StockMovementResponse{}, err
	}

	// convert movement to dto
	movementDTO := helper.ConvertToDTOStockMovementSingle(newMovement)
	return movementDTO, nil
}

// kartu stock
func (s *stockMovementService) GetStockCard(ctx context.Context, itemVariantID string) ([]dto.StockCardResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get movements
	movements, err := s.stockMovementRepo.GetMovementsByVariant(ctx, tenantID, itemVariantID)

	if err != nil {
		return nil, err
	}

	// build stock card
	var stockCards []dto.StockCardResponse

	var runningBalance float64

	for _, m := range movements {
		var qtyIn float64
		var qtyOut float64

		// in / out
		if m.Qty > 0 {
			qtyIn = m.Qty
		} else {
			qtyOut = m.Qty * -1
		}

		// running balance, saldo stok item per baris dikalkulasi terus
		runningBalance += m.Qty

		// append response
		stockCards = append(stockCards, dto.StockCardResponse{
			MovementDate:  m.CreatedAt,
			MovementType:  m.MovementType,
			QtyIn:         qtyIn,
			QtyOut:        qtyOut,
			Balance:       runningBalance,
			ReferenceType: m.ReferenceType,
			ReferenceID:   m.ReferenceID,
			Notes:         m.Notes,
			CreatedByName: m.CreatedByUser.FullName,
		})
	}
	return stockCards, nil
}
