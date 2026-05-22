package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

/*
	Gambaran besar proses create sales, sale items, dan stock movement
	tx := db.Begin()
	saleRepo.CreateSale(tx)
	saleItemRepo.CreateSaleItem(tx)
	stockRepo.CreateMovement(tx)
	tx.Commit()
	Penjelasan:
		-atomic
		-konsisten
		-aman rollback
		-tidak corrupt
*/

// interface
type SaleRepository interface {
	CreateSale(ctx context.Context, tx *gorm.DB, sale *model.Sale) error
	GetSaleByID(ctx context.Context, id string) (*model.Sale, error)
}

// struct implementasi
type saleRepository struct {
	db *gorm.DB
}

// constructor
func NewSaleRepository(db *gorm.DB) SaleRepository {
	return &saleRepository{
		db: db,
	}
}

// struct method.
// CreateSale pakai tx bukan r.db karena sale, sale item, dan stock movement harus dalam 1 transaction yang sama
func (r *saleRepository) CreateSale(ctx context.Context, tx *gorm.DB, sale *model.Sale) error {
	return tx.WithContext(ctx).Create(sale).Error
}

func (r *saleRepository) GetSaleByID(ctx context.Context, id string) (*model.Sale, error) {
	var sale model.Sale
	err := r.db.WithContext(ctx).Preload("Tenant").Preload("Cashier").First(&sale, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}
