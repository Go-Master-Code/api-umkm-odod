package service

import (
	"context"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"
)

// interface
type CatalogItemService interface {
	GetCatalogItems(ctx context.Context, name string) ([]dto.CatalogItemResponse, error)
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
	ci, err := s.repo.GetCatalogItems(ctx, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	ciDTO := helper.ConvertToDTOCatalogItemPlural(ci)
	return ciDTO, nil
}
