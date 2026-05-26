package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type StockMovementRepository interface {
	CreateMovement(ctx context.Context, tx *gorm.DB, movement *model.StockMovement) error
	GetMovementByID(ctx context.Context, id string) (*model.StockMovement, error)
	GetMovementsByVariant(ctx context.Context, itemVariantID string) ([]model.StockMovement, error)
	GetCurrentStock(ctx context.Context, tx *gorm.DB, itemVariantID string) (float64, error)
}

// struct implementasi
type stockMovementRepository struct {
	db *gorm.DB
}

// constructor
func NewStockMovementRepository(db *gorm.DB) StockMovementRepository {
	return &stockMovementRepository{
		db: db,
	}
}

// struct method
func (r *stockMovementRepository) CreateMovement(ctx context.Context, tx *gorm.DB, movement *model.StockMovement) error {
	return r.db.WithContext(ctx).Create(movement).Error
}

func (r *stockMovementRepository) GetMovementByID(ctx context.Context, id string) (*model.StockMovement, error) {
	var sm model.StockMovement
	err := r.db.WithContext(ctx).Preload("Tenant").Preload("ItemVariant").Preload("CreatedByUser").First(&sm, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &sm, nil
}

func (r *stockMovementRepository) GetMovementsByVariant(ctx context.Context, itemVariantID string) ([]model.StockMovement, error) {
	var movements []model.StockMovement
	err := r.db.WithContext(ctx).
		Preload("Tenant").
		Preload("ItemVariant").
		Preload("CreatedByUser").
		Where("item_variant_id = ?", itemVariantID).
		Order("created_at DESC").
		Find(&movements).Error

	if err != nil {
		return nil, err
	}

	return movements, nil
}

func (r *stockMovementRepository) GetCurrentStock(ctx context.Context, tx *gorm.DB, itemVariantID string) (float64, error) {
	var totalStock float64

	err := r.db.
		WithContext(ctx).
		Model(&model.StockMovement{}).
		Where("item_variant_id = ?", itemVariantID).
		Select("COALESCE(SUM(qty), 0)").
		Scan(&totalStock).
		Error

	// coalesce -> jika belum ada movement sum(qty) maka akan return null bukan 0

	if err != nil {
		return 0, err
	}

	return totalStock, nil
}
