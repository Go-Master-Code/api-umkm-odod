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
type CatalogCategoryService interface {
	GetCatalogCategories(ctx context.Context, name string) ([]dto.CatalogCategoryResponse, error)
	GetCatalogCategoryByID(ctx context.Context, id string) (dto.CatalogCategoryResponse, error)
	CreateCatalogCategory(ctx context.Context, req dto.CreateCatalogCategoryRequest) (dto.CatalogCategoryResponse, error)
	UpdateCatalogCategory(ctx context.Context, id string, req dto.UpdateCatalogCategoryRequest) (dto.CatalogCategoryResponse, error)
	DeleteCatalogCategory(ctx context.Context, id string) (dto.CatalogCategoryResponse, error)
}

// struct implementasi
type catalogCategoryService struct {
	repo repository.CatalogCategoryRepository
}

// constructor
func NewCatalogCategoryService(repo repository.CatalogCategoryRepository) CatalogCategoryService {
	return &catalogCategoryService{
		repo: repo,
	}
}

// struct method
func (s *catalogCategoryService) GetCatalogCategories(ctx context.Context, name string) ([]dto.CatalogCategoryResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	cc, err := s.repo.GetCatalogCategories(ctx, tenantID, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	ccDTO := helper.ConvertToDTOCatalogCategoryPlural(cc)
	return ccDTO, nil
}

func (s *catalogCategoryService) GetCatalogCategoryByID(ctx context.Context, id string) (dto.CatalogCategoryResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	cc, err := s.repo.GetCatalogCategoryByID(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogCategoryResponse{}, err
	}

	// convert model to dto
	ccDTO := helper.ConvertToDTOCatalogCategorySingle(cc)
	return ccDTO, nil
}

func (s *catalogCategoryService) CreateCatalogCategory(ctx context.Context, req dto.CreateCatalogCategoryRequest) (dto.CatalogCategoryResponse, error) {
	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// parsing req dto to model
	cc := model.CatalogCategory{
		ID:       uuid.NewString(),
		Name:     req.Name,
		TenantID: tenantID,
	}

	err := s.repo.CreateCatalogCategory(ctx, &cc)

	if err != nil {
		return dto.CatalogCategoryResponse{}, err
	}

	// get data by id untuk ditampilkan di response
	newCC, err := s.repo.GetCatalogCategoryByID(ctx, tenantID, cc.ID)

	// convert model to dto
	ccDTO := helper.ConvertToDTOCatalogCategorySingle(newCC)
	return ccDTO, nil
}

func (s *catalogCategoryService) UpdateCatalogCategory(ctx context.Context, id string, req dto.UpdateCatalogCategoryRequest) (dto.CatalogCategoryResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// mapping request ke map
	var updateMap = map[string]any{}

	if req.Name != nil {
		updateMap["name"] = req.Name
	}

	err := s.repo.UpdateCatalogCategory(ctx, tenantID, id, updateMap)
	if err != nil {
		return dto.CatalogCategoryResponse{}, err
	}

	// get data by id untuk melihat perubahan
	cc, err := s.repo.GetCatalogCategoryByID(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogCategoryResponse{}, err
	}

	// convert model to dto
	ccDTO := helper.ConvertToDTOCatalogCategorySingle(cc)
	return ccDTO, nil
}

func (s catalogCategoryService) DeleteCatalogCategory(ctx context.Context, id string) (dto.CatalogCategoryResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data by id dulu untuk response
	cc, err := s.repo.GetCatalogCategoryByID(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogCategoryResponse{}, err
	}

	err = s.repo.DeleteCatalogCategory(ctx, tenantID, id)
	if err != nil {
		return dto.CatalogCategoryResponse{}, err
	}

	// convert model to dto
	ccDTO := helper.ConvertToDTOCatalogCategorySingle(cc)
	return ccDTO, nil
}
