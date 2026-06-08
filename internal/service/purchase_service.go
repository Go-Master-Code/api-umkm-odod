package service

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"

	"gorm.io/gorm"
)

/*
	BEGIN TX
	↓
	validasi supplier
	↓
	create purchase header
	↓
	loop purchase items
		↓
		get item variant
		↓
		create purchase item
		↓
		create stock movement (+qty)
	↓
	hitung subtotal
	↓
	hitung tax
	↓
	hitung grand total
	↓
	update purchase header
	↓
	COMMIT
	↓
	GetPurchaseByID()
	↓
	convert DTO
*/

// interface
type PurchaseService interface {
	CreateSale(ctx context.Context, req dto.CreatePurchaseRequest) (dto.PurchaseResponse, error)
}

// struct implementasi
type purchaseService struct {
	db                *gorm.DB
	purchaseRepo      repository.PurchaseRepository
	purchaseItemRepo  repository.PurchaseItemRepository
	itemVariantRepo   repository.ItemVariantRepository
	stockMovementRepo repository.StockMovementRepository
}

// constructor
func NewPurchaseService(db *gorm.DB, purchaseRepo repository.PurchaseRepository, purchaseItemRepo repository.PurchaseItemRepository, itemVariantRepo repository.ItemVariantRepository, stockMovementRepo repository.StockMovementRepository) PurchaseService {
	return &purchaseService{
		db:                db,
		purchaseRepo:      purchaseRepo,
		purchaseItemRepo:  purchaseItemRepo,
		itemVariantRepo:   itemVariantRepo,
		stockMovementRepo: stockMovementRepo,
	}
}

// struct method
func (s *purchaseService) CreateSale(ctx context.Context, req dto.CreatePurchaseRequest) (dto.PurchaseResponse, error) {

}
