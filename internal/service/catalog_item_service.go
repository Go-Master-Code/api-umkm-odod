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
type CatalogItemService interface {
	GetCatalogItems(ctx context.Context, name string) ([]dto.CatalogItemResponse, error)
	GetCatalogItemByID(ctx context.Context, id string) (dto.CatalogItemResponse, error)
	CreateCatalogItem(ctx context.Context, req dto.CreateCatalogItemRequest) (dto.CatalogItemResponse, error)
	UpdateCatalogItem(ctx context.Context, id string, req dto.UpdateCatalogItemRequest) (dto.CatalogItemResponse, error)
	DeleteCatalogItem(ctx context.Context, id string) (dto.CatalogItemResponse, error)
}

// struct implementasi
type catalogItemService struct {
	repo repository.CatalogItemRepository
}

// constructor
func NewCatalogItemService(repo repository.CatalogItemRepository) CatalogItemService {
	return &catalogItemService{
		repo: repo,
	}
}

// struct method
func (s *catalogItemService) GetCatalogItems(ctx context.Context, name string) ([]dto.CatalogItemResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	ci, err := s.repo.GetCatalogItems(ctx, tenantID, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	ciDTO := helper.ConvertToDTOCatalogItemPlural(ci)
	return ciDTO, nil
}

func (s *catalogItemService) GetCatalogItemByID(ctx context.Context, id string) (dto.CatalogItemResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	ci, err := s.repo.GetCatalogItemByID(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogItemResponse{}, err
	}

	// convert model to dto
	ciDTO := helper.ConvertToDTOCatalogItemSingle(ci)
	return ciDTO, nil
}

func (s *catalogItemService) CreateCatalogItem(ctx context.Context, req dto.CreateCatalogItemRequest) (dto.CatalogItemResponse, error) {
	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// parsing req dto to model
	ci := model.CatalogItem{
		ID:          uuid.NewString(),
		Name:        req.Name,
		TenantID:    tenantID,
		CategoryID:  req.CategoryID,
		Description: req.Description,
		IsActive:    true,
	}

	err := s.repo.CreateCatalogItem(ctx, &ci)

	if err != nil {
		return dto.CatalogItemResponse{}, err
	}

	// get data by id untuk ditampilkan di response
	newCC, err := s.repo.GetCatalogItemByID(ctx, tenantID, ci.ID)

	// convert model to dto
	ciDTO := helper.ConvertToDTOCatalogItemSingle(newCC)
	return ciDTO, nil
}

func (s *catalogItemService) UpdateCatalogItem(ctx context.Context, id string, req dto.UpdateCatalogItemRequest) (dto.CatalogItemResponse, error) {
	// mapping request ke map
	var updateMap = map[string]any{}

	if req.Name != nil {
		updateMap["name"] = req.Name
	}
	if req.CategoryID != nil {
		updateMap["category_id"] = req.CategoryID
	}
	if req.Description != nil {
		updateMap["description"] = req.Description
	}
	if req.IsActive != nil {
		updateMap["is_active"] = req.IsActive
	}

	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	err := s.repo.UpdateCatalogItem(ctx, tenantID, id, updateMap)
	if err != nil {
		return dto.CatalogItemResponse{}, err
	}

	// get data by id untuk melihat perubahan
	ci, err := s.repo.GetCatalogItemByID(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogItemResponse{}, err
	}

	// convert model to dto
	ciDTO := helper.ConvertToDTOCatalogItemSingle(ci)
	return ciDTO, nil
}

func (s catalogItemService) DeleteCatalogItem(ctx context.Context, id string) (dto.CatalogItemResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data by id dulu untuk response
	ci, err := s.repo.GetCatalogItemByID(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogItemResponse{}, err
	}

	err = s.repo.DeleteCatalogItem(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogItemResponse{}, err
	}

	// convert model to dto
	ciDTO := helper.ConvertToDTOCatalogItemSingle(ci)
	return ciDTO, nil
}
