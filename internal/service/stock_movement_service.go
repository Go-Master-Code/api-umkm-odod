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
)

// interface
type StockMovementService interface {
	AddStock(ctx context.Context, req dto.AddStockRequest) (dto.StockMovementResponse, error)
	ReduceStock(ctx context.Context, req dto.ReduceStockRequest) (dto.StockMovementResponse, error)
	GetCurrentStock(ctx context.Context, itemVariantID string) (dto.CurrentStockResponse, error)
}

// struct implementasi
type stockMovementService struct {
	repo repository.StockMovementRepository
}

// constructor
func NewStockMovementService(repo repository.StockMovementRepository) StockMovementService {
	return &stockMovementService{
		repo: repo,
	}
}

// struct method
func (s *stockMovementService) AddStock(ctx context.Context, req dto.AddStockRequest) (dto.StockMovementResponse, error) {
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

	err := s.repo.CreateMovement(ctx, &movement)
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
	currentStock, err := s.repo.GetCurrentStock(ctx, req.ItemVariantID)
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
	err = s.repo.CreateMovement(ctx, &movement)

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
	// code
	stock, err := s.repo.GetCurrentStock(ctx, itemVariantID)
	if err != nil {
		return dto.CurrentStockResponse{}, err
	}

	// convert to dto CurrentStockResponse
	currentStockDTO := helper.ConvertToDTOCurrentStock(itemVariantID, stock)
	return currentStockDTO, nil
}
