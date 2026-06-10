package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type PurchaseReturnRepository interface {
	CreatePurchaseReturn(ctx context.Context, tx *gorm.DB, purchaseReturn *model.PurchaseReturn) error
	GetPurchaseReturnByID(ctx context.Context, tenantID string, id string) (*model.PurchaseReturn, error)
}

// struct implementasi
type purchaseReturnRepository struct {
	db *gorm.DB
}

// constructor
func NewPurchaseReturnRepository(db *gorm.DB) PurchaseReturnRepository {
	return &purchaseReturnRepository{db: db}
}

// struct method
func (r *purchaseReturnRepository) CreatePurchaseReturn(ctx context.Context, tx *gorm.DB, purchaseReturn *model.PurchaseReturn) error {
	return tx.WithContext(ctx).Create(purchaseReturn).Error
}

func (r *purchaseReturnRepository) GetPurchaseReturnByID(ctx context.Context, tenantID string, id string) (*model.PurchaseReturn, error) {
	var purchaseReturn model.PurchaseReturn
	// preload Tenant, Purchase, User
	err := r.db.WithContext(ctx).
		Preload("Tenant").
		Preload("Purchase").
		Preload("User").
		Preload("PurchaseReturnItems").
		Preload("PurchaseReturnItems.Tenant").         // preload nested relation dari purchase return item
		Preload("PurchaseReturnItems.PurchaseReturn"). // preload nested relation dari purchase return item
		Preload("PurchaseReturnItems.ItemVariant").    // preload nested relation dari purchase return item
		Where("purchase_returns.tenant_id = ? AND purchase_returns.id = ?", tenantID, id).
		First(&purchaseReturn).Error

	if err != nil {
		return nil, err
	}

	return &purchaseReturn, nil
}
