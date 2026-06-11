package service

import (
	"context"
	"fmt"
	"time"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/repository"

	"github.com/google/uuid"
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
	GetAllPurchases(ctx context.Context, query dto.GetAllPurchasesQuery) ([]dto.PurchaseResponse, int64, error)
	CreatePurchase(ctx context.Context, req dto.CreatePurchaseRequest) (dto.PurchaseResponse, error)
	GetPurchaseByID(ctx context.Context, id string) (dto.PurchaseResponse, error)
}

// struct implementasi
type purchaseService struct {
	db                *gorm.DB
	purchaseRepo      repository.PurchaseRepository
	purchaseItemRepo  repository.PurchaseItemRepository
	itemVariantRepo   repository.ItemVariantRepository
	stockMovementRepo repository.StockMovementRepository
	supplierRepo      repository.SupplierRepository
}

// constructor
func NewPurchaseService(db *gorm.DB, purchaseRepo repository.PurchaseRepository, purchaseItemRepo repository.PurchaseItemRepository, itemVariantRepo repository.ItemVariantRepository, stockMovementRepo repository.StockMovementRepository, supplierRepo repository.SupplierRepository) PurchaseService {
	return &purchaseService{
		db:                db,
		purchaseRepo:      purchaseRepo,
		purchaseItemRepo:  purchaseItemRepo,
		itemVariantRepo:   itemVariantRepo,
		stockMovementRepo: stockMovementRepo,
		supplierRepo:      supplierRepo,
	}
}

// struct method
func (s *purchaseService) GetAllPurchases(ctx context.Context, query dto.GetAllPurchasesQuery) ([]dto.PurchaseResponse, int64, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	// get data from repository
	purchases, total, err := s.purchaseRepo.GetAllPurchases(ctx, tenantID, query)
	if err != nil {
		return nil, 0, err
	}
	// convert model to dto
	purchasesDTO := helper.ConvertToDTOPurchasePlural(purchases)
	// jika semua sukses
	return purchasesDTO, total, nil
}

func (s *purchaseService) CreatePurchase(ctx context.Context, req dto.CreatePurchaseRequest) (dto.PurchaseResponse, error) {
	// begin database transaction
	tx := s.db.Begin()

	if tx.Error != nil {
		return dto.PurchaseResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	// ========================================
	// AMBIL DATA TENANT DARI CONTEXT JWT
	// ========================================
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// validasi supplier
	supplier, err := s.supplierRepo.GetSupplierByID(ctx, tenantID, req.SupplierID)
	if err != nil {
		tx.Rollback()
		return dto.PurchaseResponse{}, err
	}

	// ambil data user dari context
	userID := ctx.Value(constants.ContextUserID).(string)

	// ========================================
	// GENERATE PURCHASE NUMBER
	// ========================================
	purchaseNumber := fmt.Sprintf(
		"PO-%d",
		time.Now().Unix(),
	)

	// ========================================
	// CREATE PURCHASE HEADER -> master purchase
	// ========================================
	purchase := model.Purchase{
		ID:             uuid.NewString(),
		TenantID:       tenantID,
		PurchaseNumber: purchaseNumber,
		SupplierID:     supplier.ID, // ambil dari hasil validasi supplier di atas
		InvoiceNumber:  req.InvoiceNumber,
		DiscountAmount: req.DiscountAmount, // discount di master purchase
		Notes:          req.Notes,
		CreatedBy:      userID,
	}

	// simpan master purchase
	err = s.purchaseRepo.CreatePurchase(ctx, tx, &purchase)
	if err != nil {
		tx.Rollback() // jika terjadi error saat insert master purchase, rollback
		return dto.PurchaseResponse{}, err
	}

	// inisiasi grandTotal
	var grandTotal float64

	// loop purchase items
	for _, item := range req.Items {
		// ========================================
		// AMBIL ITEM VARIANT DARI DATABASE
		// harga harus trusted dari DB
		// ========================================
		variant, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, item.ItemVariantID)
		if err != nil {
			tx.Rollback()
			return dto.PurchaseResponse{}, err
		}

		// hitung subtotal per row purchase item
		subtotal := (item.Qty * item.CostPrice) - item.DiscountAmount

		// ========================================
		// CREATE PURCHASE ITEM
		// gunakan snapshot agar histori immutable
		// ========================================

		purchaseItem := model.PurchaseItem{
			ID:                  uuid.NewString(),
			TenantID:            tenantID,
			PurchaseID:          purchase.ID,
			ItemVariantID:       item.ItemVariantID,
			ItemNameSnapshot:    variant.Item.Name,
			VariantNameSnapshot: variant.VariantName,
			SKUSnapshot:         variant.SKU,
			Qty:                 item.Qty,
			CostPrice:           item.CostPrice,
			DiscountAmount:      item.DiscountAmount,
			Subtotal:            subtotal,
		}

		// simpan detil purchase ke db
		err = s.purchaseItemRepo.CreatePurchaseItem(ctx, tx, &purchaseItem)
		if err != nil {
			tx.Rollback()
			return dto.PurchaseResponse{}, err
		}

		// ========================================
		// CREATE STOCK MOVEMENT
		// stok masuk, qty selalu positif
		// ========================================
		movement := model.StockMovement{
			ID:            uuid.NewString(),
			TenantID:      tenantID,
			ItemVariantID: variant.ID,
			MovementType:  constants.MovementPurchase,
			Qty:           item.Qty, // positif karena barang masuk
			ReferenceType: "PURCHASE",
			ReferenceID:   purchase.ID,
			Notes:         "purchase transaction",
			CreatedBy:     userID,
		}

		// simpan stock movement
		err = s.stockMovementRepo.CreateMovement(ctx, tx, &movement)
		if err != nil {
			tx.Rollback()
			return dto.PurchaseResponse{}, err
		}

		// akumulasi grand total
		grandTotal += subtotal

		// increment master sale subtotal
		purchase.Subtotal += subtotal // subtotal = subtotal per row
	}

	// hitung pajak (jika ada)
	purchase.TaxAmount = grandTotal / 10

	// hitung grand total master purchase
	purchase.GrandTotal = purchase.Subtotal - purchase.DiscountAmount + purchase.TaxAmount

	// update subtotal dan grand total ke master purchase pakai map
	// subtotal = total agregat dari masing-masing row item (harga * qty) tiap row
	err = tx.
		WithContext(ctx).
		Model(&purchase).
		Updates(map[string]any{
			"subtotal":    purchase.Subtotal,
			"tax_amount":  purchase.TaxAmount,
			"grand_total": purchase.GrandTotal,
		}).Error

	if err != nil {
		tx.Rollback()
		return dto.PurchaseResponse{}, err
	}

	// commit transaction
	err = tx.Commit().Error

	if err != nil {
		return dto.PurchaseResponse{}, err
	}

	// GetPurchaseByID untuk dapat data purchase yang sudah terbaru dan preload semua relasi
	newPurchase, err := s.purchaseRepo.GetPurchaseByID(ctx, tenantID, purchase.ID)

	if err != nil {
		return dto.PurchaseResponse{}, err
	}

	// Convert model to dto
	purchaseDTO := helper.ConvertToDTOPurchaseSingle(newPurchase)

	return purchaseDTO, nil
}

func (s *purchaseService) GetPurchaseByID(ctx context.Context, id string) (dto.PurchaseResponse, error) {
	// get tenant id from context
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	// akses repo
	purchase, err := s.purchaseRepo.GetPurchaseByID(ctx, tenantID, id)
	if err != nil {
		return dto.PurchaseResponse{}, err
	}

	// convert model to dto
	purchaseDTO := helper.ConvertToDTOPurchaseSingle(purchase)

	return purchaseDTO, nil
}
