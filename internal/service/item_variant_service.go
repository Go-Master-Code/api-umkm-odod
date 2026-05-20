package service

import (
	"context"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"
)

// interface
type ItemVariantService interface {
	GetItemVariants(ctx context.Context, name string) ([]dto.ItemVariantResponse, error)
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
