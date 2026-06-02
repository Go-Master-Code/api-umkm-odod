package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ItemVariantRepository interface {
	GetAllItemVariants(ctx context.Context, tenantID string) ([]model.ItemVariant, error)
	GetItemVariants(ctx context.Context, tenantID string, name string) ([]model.ItemVariant, error)
	GetItemVariantByID(ctx context.Context, tenantID string, id string) (*model.ItemVariant, error)
	CreateItemVariant(ctx context.Context, iv *model.ItemVariant) error
	UpdateItemVariant(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	DeleteItemVariant(ctx context.Context, tenantID string, id string) error
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
func (r *itemVariantRepository) GetAllItemVariants(ctx context.Context, tenantID string) ([]model.ItemVariant, error) {
	var variants []model.ItemVariant
	err := r.db.WithContext(ctx).Preload("Item").Where("tenant_id = ?", tenantID).Find(&variants).Error
	if err != nil {
		return nil, err
	}

	return variants, nil
}

func (r *itemVariantRepository) GetItemVariants(ctx context.Context, tenantID string, name string) ([]model.ItemVariant, error) {
	var iv []model.ItemVariant
	// query default
	query := r.db.WithContext(ctx).Preload("Tenant").Preload("Item").Where("tenant_id = ?", tenantID) // Preload Item sesuaikan dengan model item_variants.go

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

func (r *itemVariantRepository) GetItemVariantByID(ctx context.Context, tenantID string, id string) (*model.ItemVariant, error) {
	var iv model.ItemVariant
	err := r.db.WithContext(ctx).Preload("Tenant").Preload("Item").Where("tenant_id = ?", tenantID).First(&iv, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &iv, nil
}

func (r *itemVariantRepository) CreateItemVariant(ctx context.Context, iv *model.ItemVariant) error {
	return r.db.WithContext(ctx).Create(iv).Error
}

func (r *itemVariantRepository) UpdateItemVariant(ctx context.Context, tenantID string, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.ItemVariant{}).Where("id = ? AND tenant_id = ?", id, tenantID).Updates(updateMap).Error
}

func (r *itemVariantRepository) DeleteItemVariant(ctx context.Context, tenantID string, id string) error {
	return r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&model.ItemVariant{}).Error
}
