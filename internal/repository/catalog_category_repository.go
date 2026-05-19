package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type CatalogCategoryRepository interface {
	GetCatalogCategories(ctx context.Context, name string) ([]model.CatalogCategory, error)
}

// struct implementasi
type catalogCategoryRepository struct {
	db *gorm.DB
}

// constructor
func NewCatalogCategoryRepository(db *gorm.DB) CatalogCategoryRepository {
	return &catalogCategoryRepository{
		db: db,
	}
}

// struct method
func (r *catalogCategoryRepository) GetCatalogCategories(ctx context.Context, name string) ([]model.CatalogCategory, error) {
	var cc []model.CatalogCategory
	// query utama
	query := r.db.WithContext(ctx).Preload("Tenant")

	// jika name tidak kosong
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	err := query.Find(&cc).Error
	if err != nil {
		return nil, err
	}

	return cc, nil
}
