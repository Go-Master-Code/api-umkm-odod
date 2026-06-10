package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type PurchaseReturnItemRepository interface {
	CreatePurchaseReturnItem(ctx context.Context, tx *gorm.DB, purchaseReturnItem *model.PurchaseReturnItem) error
	/*
		CreatePurchaseReturnItem harus menggunakan tx
		karena purchase return header, purchase return item,
		dan stock movement harus berada dalam
		transaction yang sama.
	*/
}

// struct implementasi
type purchaseReturnItemRepository struct {
	db *gorm.DB
}

// constructur
func NewPurchaseReturnItemRepository(db *gorm.DB) PurchaseReturnItemRepository {
	return &purchaseReturnItemRepository{
		db: db,
	}
}

// struct method
func (r *purchaseReturnItemRepository) CreatePurchaseReturnItem(ctx context.Context, tx *gorm.DB, purchaseReturnItem *model.PurchaseReturnItem) error {
	return tx.WithContext(ctx).Create(purchaseReturnItem).Error
}
