package service

import (
	"context"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/repository"

	"github.com/google/uuid"
)

// interface
type ItemVariantService interface {
	GetItemVariants(ctx context.Context, name string) ([]dto.ItemVariantResponse, error)
	GetItemVariantByID(ctx context.Context, id string) (dto.ItemVariantResponse, error)
	GetLowStockItem(ctx context.Context) ([]dto.LowStockResponse, error)
	CreateItemVariant(ctx context.Context, req dto.CreateItemVariantRequest) (dto.ItemVariantResponse, error)
	UpdateItemVariant(ctx context.Context, id string, req dto.UpdateItemVariantRequest) (dto.ItemVariantResponse, error)
	DeleteItemVariant(ctx context.Context, id string) (dto.ItemVariantResponse, error)
}

// struct implementasi
type itemVariantService struct {
	itemVariantRepo   repository.ItemVariantRepository
	stockMovementRepo repository.StockMovementRepository
}

// constructor
func NewItemVariantService(itemVariantRepo repository.ItemVariantRepository, stockMovementRepo repository.StockMovementRepository) ItemVariantService {
	return &itemVariantService{
		itemVariantRepo:   itemVariantRepo,
		stockMovementRepo: stockMovementRepo,
	}
}

// struct method
func (s *itemVariantService) GetItemVariants(ctx context.Context, name string) ([]dto.ItemVariantResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	iv, err := s.itemVariantRepo.GetItemVariants(ctx, tenantID, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantPlural(iv)
	return ivDTO, nil
}

func (s *itemVariantService) GetItemVariantByID(ctx context.Context, id string) (dto.ItemVariantResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	iv, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(iv)
	return ivDTO, err
}

func (s *itemVariantService) GetLowStockItem(ctx context.Context) ([]dto.LowStockResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	variants, err := s.itemVariantRepo.GetAllItemVariants(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	var lowStock []dto.LowStockResponse
	for _, variant := range variants {
		// ambil stok realtime
		// kalau tx ada -> pakai transaction
		// kalau tx nil -> pakai repository db biasa, dan query akan menggunakan r.db, tidak pakai mode transaction (tx)
		currentStock, err := s.stockMovementRepo.GetCurrentStock(ctx, tenantID, nil, variant.ID)

		if err != nil {
			return nil, err
		}

		// cek low stock -> hanya record data ini yang akan muncul
		if currentStock <= variant.MinimumStock {
			lowStock = append(lowStock, dto.LowStockResponse{
				ItemVariantID: variant.ID,
				ItemName:      variant.Item.Name,
				SKU:           variant.SKU,
				VariantName:   variant.VariantName,
				CurrentStock:  currentStock,
				MinimumStock:  variant.MinimumStock,
			})
		}
	}
	return lowStock, nil
}

func (s *itemVariantService) CreateItemVariant(ctx context.Context, req dto.CreateItemVariantRequest) (dto.ItemVariantResponse, error) {
	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// parsing dto ke model
	iv := model.ItemVariant{
		ID:          uuid.NewString(),
		TenantID:    tenantID,
		ItemID:      req.ItemID,
		SKU:         req.SKU,
		Barcode:     req.Barcode,
		CostPrice:   req.CostPrice,
		VariantName: req.VariantName,
		IsActive:    req.IsActive,
	}

	err := s.itemVariantRepo.CreateItemVariant(ctx, &iv)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// get data by id
	newIV, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, iv.ID)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(newIV)
	return ivDTO, nil
}

func (s *itemVariantService) UpdateItemVariant(ctx context.Context, id string, req dto.UpdateItemVariantRequest) (dto.ItemVariantResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// mapping req ke map
	var updateMap = map[string]any{}

	if req.ItemID != nil {
		updateMap["item_id"] = req.ItemID
	}
	if req.SKU != nil {
		updateMap["sku"] = req.SKU
	}
	if req.Barcode != nil {
		updateMap["barcode"] = req.Barcode
	}
	if req.VariantName != nil {
		updateMap["variant_name"] = req.VariantName
	}
	if req.CostPrice != nil {
		updateMap["cost_price"] = req.CostPrice
	}
	if req.IsActive != nil {
		updateMap["is_active"] = req.IsActive
	}

	err := s.itemVariantRepo.UpdateItemVariant(ctx, tenantID, id, updateMap)

	// get data by id
	iv, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(iv)

	return ivDTO, nil
}

func (s *itemVariantService) DeleteItemVariant(ctx context.Context, id string) (dto.ItemVariantResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get id dulu untuk response api
	iv, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// delete data
	err = s.itemVariantRepo.DeleteItemVariant(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(iv)
	return ivDTO, nil
}
