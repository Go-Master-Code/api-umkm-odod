package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

/*
	Gambaran besar proses create purchase, purchase items, dan stock movement
	tx := db.Begin()
	purchaseRepo.CreatePurchase(tx)
	purchaseItemRepo.CreatePurchaseItem(tx)
	stockRepo.CreateMovement(tx)
	tx.Commit()
	Penjelasan:
		-atomic
		-konsisten
		-aman rollback
		-tidak corrupt
*/

// interface
type PurchaseRepository interface {
	GetAllPurchases(ctx context.Context, tenantID string, query dto.GetAllPurchasesQuery) ([]model.Purchase, int64, error)
	CreatePurchase(ctx context.Context, tx *gorm.DB, sale *model.Purchase) error
	GetPurchaseByID(ctx context.Context, tenantID string, id string) (*model.Purchase, error) // perlu tenant isolation agar tenant A tidak bisa akses invoice tenant B
}

// struct implementasi
type purchaseRepository struct {
	db *gorm.DB
}

// constructor
func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{
		db: db,
	}
}

// struct method.

func (r *saleRepository) GetAllPurchases(ctx context.Context, tenantID string, query dto.GetAllPurchasesQuery) ([]model.Purchase, int64, error) {
	var purchases []model.Purchase
	var total int64

	offset := (query.Page - 1) * query.Limit

	// base query
	baseQuery := r.db.WithContext(ctx).Model(&model.Purchase{}).Where("tenant_id = ?", tenantID)

	// // filter payment status
	// if query.PaymentStatus != "" {
	// 	baseQuery = baseQuery.Where("payment_status = ?", query.PaymentStatus)
	// }

	// search
	if query.Search != "" {
		search := "%" + query.Search + "%"
		baseQuery = baseQuery.Where("invoice_number LIKE ? OR purchase_number LIKE ?", search, search) // query untuk cari data yang invoice atau purchase number nya like ...
	}

	// count total row(s) found -> diperlukan untuk frontend
	err := baseQuery.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// get data
	err = baseQuery.Preload("Tenant").Preload("User").Preload("Purchase").Preload("PurchaseItems.Tenant").Preload("PurchaseItems.ItemVariant").Order("created_at DESC").Limit(query.Limit).Offset(offset).Find(&purchases).Error
	if err != nil {
		return nil, 0, err
	}

	// jika semua sukses
	return purchases, total, nil
}

// CreateSale pakai tx bukan r.db karena sale, sale item, dan stock movement harus dalam 1 transaction yang sama
func (r *saleRepository) CreateSale(ctx context.Context, tx *gorm.DB, sale *model.Sale) error {
	return tx.WithContext(ctx).Create(sale).Error
}

func (r *saleRepository) GetSaleByID(ctx context.Context, tenantID string, id string) (*model.Sale, error) {
	var sale model.Sale
	err := r.db.
		WithContext(ctx).
		Preload("Tenant").
		Preload("Cashier").
		Preload("SaleItems").
		Preload("SaleItems.Tenant").      // preload nested relation dari sale item
		Preload("SaleItems.Sale").        // preload nested relation dari sale item
		Preload("SaleItems.ItemVariant"). // preload nested relation dari sale item
		Where("id = ? AND tenant_id = ?", id, tenantID).
		First(&sale).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}
