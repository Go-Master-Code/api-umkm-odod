package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type PurchaseItemRepository interface {
	CreatePurchaseItem(ctx context.Context, tx *gorm.DB, purchaseItem *model.PurchaseItem) error
	/*
		CreatePurchaseItem harus menggunakan tx
		karena purchase header, purchase item,
		dan stock movement harus berada dalam
		transaction yang sama.
	*/
}

// struct implementasi
type purchaseItemRepository struct {
	db *gorm.DB
}

// constructor
func NewPurchaseItemRepository(db *gorm.DB) PurchaseItemRepository {
	return &purchaseItemRepository{
		db: db,
	}
}

// struct method
func (r *purchaseItemRepository) CreatePurchaseItem(ctx context.Context, tx *gorm.DB, purchaseItem *model.PurchaseItem) error {
	return tx.WithContext(ctx).Create(purchaseItem).Error
}
