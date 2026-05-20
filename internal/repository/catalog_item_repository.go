package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type CatalogItemRepository interface {
	GetCatalogItems(ctx context.Context, name string) ([]model.CatalogItem, error)
	GetCatalogItemByID(ctx context.Context, id string) (*model.CatalogItem, error)
	CreateCatalogItem(ctx context.Context, ci *model.CatalogItem) error
	UpdateCatalogItem(ctx context.Context, id string, updateMap map[string]any) error
	DeleteCatalogItem(ctx context.Context, id string) error
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

func (r *catalogItemRepository) GetCatalogItemByID(ctx context.Context, id string) (*model.CatalogItem, error) {
	var ci model.CatalogItem
	err := r.db.WithContext(ctx).Preload("Tenant").Preload("CatalogCategory").First(&ci, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &ci, nil
}

func (r *catalogItemRepository) CreateCatalogItem(ctx context.Context, ci *model.CatalogItem) error {
	return r.db.WithContext(ctx).Create(ci).Error
}

func (r *catalogItemRepository) UpdateCatalogItem(ctx context.Context, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.CatalogItem{}).Where("id = ?", id).Updates(updateMap).Error
}

func (r *catalogItemRepository) DeleteCatalogItem(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.CatalogItem{}).Error
}
