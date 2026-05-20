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
	CreateItemVariant(ctx context.Context, req dto.CreateItemVariantRequest) (dto.ItemVariantResponse, error)
}

// struct implementasi
type itemVariantService struct {
	repo repository.ItemVariantRepository
}

// constructor
func NewItemVariantService(repo repository.ItemVariantRepository) ItemVariantService {
	return &itemVariantService{
		repo: repo,
	}
}

// struct method
func (s *itemVariantService) GetItemVariants(ctx context.Context, name string) ([]dto.ItemVariantResponse, error) {
	iv, err := s.repo.GetItemVariants(ctx, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantPlural(iv)
	return ivDTO, nil
}

func (s *itemVariantService) GetItemVariantByID(ctx context.Context, id string) (dto.ItemVariantResponse, error) {
	iv, err := s.repo.GetItemVariantByID(ctx, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(iv)
	return ivDTO, err
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

	err := s.repo.CreateItemVariant(ctx, &iv)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// get data by id
	newIV, err := s.repo.GetItemVariantByID(ctx, iv.ID)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(newIV)
	return ivDTO, nil
}
