package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ItemVariantRepository interface {
	GetItemVariants(ctx context.Context, name string) ([]model.ItemVariant, error)
	GetItemVariantByID(ctx context.Context, id string) (*model.ItemVariant, error)
	CreateItemVariant(ctx context.Context, iv *model.ItemVariant) error
	UpdateItemVariant(ctx context.Context, id string, updateMap map[string]any) error
	DeleteItemVariant(ctx context.Context, id string) error
}

// struct implementasi
type itemVariantRepository struct {
	db *gorm.DB
}

// constructor
func NewItemVariantRepository(db *gorm.DB) ItemVariantRepository {
	return &itemVariantRepository{
		db: db,
	}
}

// struct method
func (r *itemVariantRepository) GetItemVariants(ctx context.Context, name string) ([]model.ItemVariant, error) {
	var iv []model.ItemVariant
	// query default
	query := r.db.WithContext(ctx).Preload("Tenant").Preload("Item") // Preload Item sesuaikan dengan model item_variants.go

	if name != "" {
		// jika name nya tidak kosong, tambahkan query
		query = query.Where("variant_name LIKE ?", "%"+name+"%")
	}

	// find data by name
	err := query.Find(&iv).Error

	if err != nil {
		return nil, err
	}

	return iv, nil
}

func (r *itemVariantRepository) GetItemVariantByID(ctx context.Context, id string) (*model.ItemVariant, error) {
	var iv model.ItemVariant
	err := r.db.WithContext(ctx).Preload("Tenant").Preload("Item").First(&iv, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &iv, nil
}

func (r *itemVariantRepository) CreateItemVariant(ctx context.Context, iv *model.ItemVariant) error {
	return r.db.WithContext(ctx).Create(iv).Error
}

func (r *itemVariantRepository) UpdateItemVariant(ctx context.Context, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.ItemVariant{}).Where("id = ?", id).Updates(updateMap).Error
}

func (r *itemVariantRepository) DeleteItemVariant(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.ItemVariant{}).Error
}
