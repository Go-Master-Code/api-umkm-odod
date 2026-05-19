package service

import (
	"context"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"
)

// interface
type CatalogCategoryService interface {
	GetCatalogCategories(ctx context.Context, name string) ([]dto.CatalogCategoryResponse, error)
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
	cc, err := s.repo.GetCatalogCategories(ctx, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	ccDTO := helper.ConvertToDTOCatalogCategoryPlural(cc)
	return ccDTO, nil
}
