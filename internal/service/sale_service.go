package service

import (
	"context"
	"errors"
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
	===TARGET FLOW===
	BEGIN TRANSACTION
	↓
	create sale
	↓
	loop items
		↓
		get variant
		↓
		validate stock
		↓
		create sale item
		↓
		create stock movement
	↓
	update grand total
	↓
	COMMIT
*/

// interface
type SaleService interface {
	CreateSale(ctx context.Context, req dto.CreateSaleRequest) (dto.SaleResponse, error)
}

// struct implementasi
type saleService struct {
	db                *gorm.DB // ada db di service karena: transaction begin/commit/rollback dilakukan di layer service
	saleRepo          repository.SaleRepository
	saleItemRepo      repository.SaleItemRepository
	itemVariantRepo   repository.ItemVariantRepository
	stockMovementRepo repository.StockMovementRepository
}

// constructor -> ada db karena untuk transaction
func NewSaleService(
	db *gorm.DB,
	saleRepo repository.SaleRepository,
	saleItemRepo repository.SaleItemRepository,
	itemVariantRepo repository.ItemVariantRepository,
	stockMovementRepo repository.StockMovementRepository,
) SaleService {
	return &saleService{
		db:                db,
		saleRepo:          saleRepo,
		saleItemRepo:      saleItemRepo,
		itemVariantRepo:   itemVariantRepo,
		stockMovementRepo: stockMovementRepo,
	}
}

// stuct method
func (s *saleService) CreateSale(ctx context.Context, req dto.CreateSaleRequest) (dto.SaleResponse, error) {
	// ========================================
	// BEGIN DATABASE TRANSACTION
	// ========================================

	tx := s.db.Begin() // begin transaction

	if tx.Error != nil { // cek apakah gagal begin transaction
		return dto.SaleResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	// ========================================
	// AMBIL DATA USER DARI CONTEXT JWT
	// ========================================
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	userID := ctx.Value(constants.ContextUserID).(string)

	// ========================================
	// GENERATE INVOICE NUMBER
	// ========================================
	invoiceNumber := fmt.Sprintf(
		"INV-%d",
		time.Now().Unix(),
	)

	// ========================================
	// CREATE SALE HEADER -> master sale
	// ========================================
	sale := model.Sale{
		ID:             uuid.NewString(),
		TenantID:       tenantID,
		InvoiceNumber:  invoiceNumber,
		CustomerName:   req.CustomerName,
		CashierID:      userID,
		DiscountAmount: req.DiscountAmount,
		PaymentMethod:  req.PaymentMethod,
		PaymentStatus:  req.PaymentStatus,
		Notes:          req.Notes,
	}

	// simpan sale header / data master
	err := s.saleRepo.CreateSale(ctx, tx, &sale)
	if err != nil {
		tx.Rollback() // jika terjadi error saat input data header / master sale, rollback
		return dto.SaleResponse{}, err
	}

	// ========================================
	// LOOP SALE ITEMS
	// ========================================
	var grandTotal float64
	for _, item := range req.Items { // iterasi ke sale items -> lihat format dtoCreateSaleRequest
		// ========================================
		// AMBIL ITEM VARIANT DARI DATABASE
		// harga harus trusted dari DB
		// ========================================
		variant, err := s.itemVariantRepo.GetItemVariantByID(ctx, item.ItemVariantID)
		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}

		// ========================================
		// VALIDASI STOK
		// ========================================

		currentStock, err := s.stockMovementRepo.GetCurrentStock(
			ctx,
			tx,
			variant.ID,
		)

		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}

		// stok tidak boleh minus
		if currentStock < item.Qty {
			tx.Rollback()
			return dto.SaleResponse{},
				errors.New("insufficient stock")
		}

		// ========================================
		// HITUNG SUBTOTAL
		// ========================================

		subtotal := (item.Qty * variant.SellingPrice) - item.DiscountAmount

		// ========================================
		// CREATE SALE ITEM
		// gunakan snapshot agar histori immutable
		// ========================================

		saleItem := model.SaleItem{
			ID:                  uuid.NewString(),
			TenantID:            tenantID,
			SaleID:              sale.ID,
			ItemVariantID:       variant.ID,
			ItemNameSnapshot:    variant.Item.Name,
			VariantNameSnapshot: variant.VariantName,
			SKUSnapshot:         variant.SKU,
			Qty:                 item.Qty,
			UnitPrice:           variant.SellingPrice,
			DiscountAmount:      item.DiscountAmount,
			Subtotal:            subtotal,
		}

		// simpan sale item
		err = s.saleItemRepo.CreateSaleItem(ctx, tx, &saleItem)

		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}

		// ========================================
		// CREATE STOCK MOVEMENT
		// stok keluar = qty negatif
		// ========================================

		movement := model.StockMovement{
			ID:            uuid.NewString(),
			TenantID:      tenantID,
			ItemVariantID: variant.ID,
			MovementType:  constants.MovementSale,
			Qty:           -item.Qty,
			ReferenceType: "SALE",
			ReferenceID:   sale.ID,
			Notes:         "sale transaction",
			CreatedBy:     userID,
		}

		// simpan stock movement
		err = s.stockMovementRepo.CreateMovement(ctx, tx, &movement)

		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}

		// ========================================
		// TAMBAH GRAND TOTAL
		// ========================================

		grandTotal += subtotal

		// increment sale.subtotal
		sale.Subtotal += subtotal
	}

	// ========================================
	// HITUNG FINAL GRAND TOTAL
	// ========================================

	// Hitung tax, misalnya 10% dari grand total
	sale.TaxAmount = grandTotal / 10

	// update grand total setelah diskon ditambah pajak
	sale.GrandTotal = sale.Subtotal - sale.DiscountAmount + sale.TaxAmount

	// ========================================
	// UPDATE SUBTOTAL dan GRAND TOTAL KE DATABASE
	// ========================================

	// subtotal = total agregat dari masing-masing row item (harga * qty) tiap row
	err = tx.
		WithContext(ctx).
		Model(&sale).
		Update(
			"subtotal",
			sale.Subtotal,
		).Error

	if err != nil {
		tx.Rollback()
		return dto.SaleResponse{}, err
	}

	// update tax misal 10% dari grand total
	err = tx.
		WithContext(ctx).
		Model(&sale).
		Update(
			"tax_amount",
			sale.TaxAmount,
		).Error

	if err != nil {
		tx.Rollback()
		return dto.SaleResponse{}, err
	}

	// grand total = subtotal - disc (jika ada) + tax
	err = tx.
		WithContext(ctx).
		Model(&sale).
		Update(
			"grand_total",
			sale.GrandTotal,
		).Error

	if err != nil {
		tx.Rollback()
		return dto.SaleResponse{}, err
	}

	// ========================================
	// COMMIT TRANSACTION
	// ========================================

	err = tx.Commit().Error

	if err != nil {
		return dto.SaleResponse{}, err
	}

	// get data sale by id untuk preload semua relasi
	newSale, err := s.saleRepo.GetSaleByID(ctx, sale.ID)

	// ========================================
	// RESPONSE DTO
	// ========================================
	saleDTO := helper.ConvertToDTOSaleSingle(newSale)

	return saleDTO, nil
}
