package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type CatalogCategoryRepository interface {
	GetCatalogCategories(ctx context.Context, name string) ([]model.CatalogCategory, error)
	GetCatalogCategoryByID(ctx context.Context, id string) (*model.CatalogCategory, error)
	CreateCatalogCategory(ctx context.Context, cc *model.CatalogCategory) error
	UpdateCatalogCategory(ctx context.Context, id string, updateMap map[string]any) error
	DeleteCatalogCategory(ctx context.Context, id string) error
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

func (r *catalogCategoryRepository) GetCatalogCategoryByID(ctx context.Context, id string) (*model.CatalogCategory, error) {
	var cc model.CatalogCategory
	err := r.db.WithContext(ctx).Preload("Tenant").First(&cc, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &cc, nil
}

func (r *catalogCategoryRepository) CreateCatalogCategory(ctx context.Context, cc *model.CatalogCategory) error {
	return r.db.WithContext(ctx).Create(cc).Error
}

func (r *catalogCategoryRepository) UpdateCatalogCategory(ctx context.Context, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.CatalogCategory{}).Where("id = ?", id).Updates(updateMap).Error
}

func (r *catalogCategoryRepository) DeleteCatalogCategory(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.CatalogCategory{}).Error
}
