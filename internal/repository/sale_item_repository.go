package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type SaleItemRepository interface {
	CreateSaleItem(ctx context.Context, tx *gorm.DB, saleItem *model.SaleItem) error
}

// struct implementasi
type saleItemRepository struct {
	db *gorm.DB
}

// constructor
func NewSaleItemRepository(db *gorm.DB) SaleItemRepository {
	return &saleItemRepository{
		db: db,
	}
}

// struct method
func (s *saleItemRepository) CreateSaleItem(ctx context.Context, tx *gorm.DB, saleItem *model.SaleItem) error {
	return tx.WithContext(ctx).Create(saleItem).Error
}
