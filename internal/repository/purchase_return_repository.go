package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type PurchaseReturnRepository interface {
	GetAllPurchaseReturns(ctx context.Context, tenantID string, query dto.GetAllPurchaseReturnsQuery) ([]model.PurchaseReturn, int64, error)
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
func (r *purchaseReturnRepository) GetAllPurchaseReturns(ctx context.Context, tenantID string, query dto.GetAllPurchaseReturnsQuery) ([]model.PurchaseReturn, int64, error) {
	var purchaseReturns []model.PurchaseReturn
	var total int64

	// pagination sudah dilakukan di handler
	offset := (query.Page - 1) * query.Limit

	// base query
	baseQuery := r.db.WithContext(ctx).Model(&model.PurchaseReturn{}).Where("tenant_id = ?", tenantID)

	// search
	if query.Search != "" {
		search := "%" + query.Search + "%"
		baseQuery = baseQuery.Where("reason LIKE ? OR purchase_id LIKE ? OR return_number LIKE ?", search, search, search)
	}

	// count total row - diperlukan untuk frontend
	err := baseQuery.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// GET data, preload Tenant, Purchase, User
	err = baseQuery.
		Preload("Tenant").
		Preload("Purchase").
		Preload("User").
		Preload("PurchaseReturnItems").
		Preload("PurchaseReturnItems.Tenant").         // preload nested relation dari purchase return item
		Preload("PurchaseReturnItems.PurchaseReturn"). // preload nested relation dari purchase return item
		Preload("PurchaseReturnItems.ItemVariant").    // preload nested relation dari purchase return item
		Order("created_at DESC").
		Limit(query.Limit).
		Offset(offset).
		Find(&purchaseReturns).Error

	if err != nil {
		return nil, 0, err
	}

	// jika query sukses
	return purchaseReturns, total, nil
}

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
