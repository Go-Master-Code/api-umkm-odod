package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type CatalogItemRepository interface {
	GetCatalogItems(ctx context.Context, name string) ([]model.CatalogItem, error)
}

// sturct implementasi
type catalogItemRepository struct {
	db *gorm.DB
}

// constructor
func NewCatalogItemRepository(db *gorm.DB) CatalogItemRepository {
	return &catalogItemRepository{
		db: db,
	}
}

// struct method
func (r *catalogItemRepository) GetCatalogItems(ctx context.Context, name string) ([]model.CatalogItem, error) {
	var ci []model.CatalogItem
	// query default
	query := r.db.WithContext(ctx).Preload("Tenant").Preload("CatalogCategory")

	// cek name kosong atau tidak
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// lanjut find data
	err := query.Find(&ci).Error
	if err != nil {
		return nil, err
	}

	return ci, nil
}
