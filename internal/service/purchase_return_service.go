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

// interface
type PurchaseReturnService interface {
	GetAllPurchaseReturns(ctx context.Context, query dto.GetAllPurchaseReturnsQuery) ([]dto.PurchaseReturnResponse, int64, error)
	CreatePurchaseReturn(ctx context.Context, req dto.CreatePurchaseReturnRequest) (dto.PurchaseReturnResponse, error)
	GetPurchaseReturnByID(ctx context.Context, id string) (dto.PurchaseReturnResponse, error)
}

// struct implementasi
type purchaseReturnService struct {
	db                     *gorm.DB
	purchaseReturnRepo     repository.PurchaseReturnRepository
	purchaseReturnItemRepo repository.PurchaseReturnItemRepository
	itemVariantRepo        repository.ItemVariantRepository
	stockMovementRepo      repository.StockMovementRepository

	// log
	activityLogService ActivityLogService // jangan pakai package service, karena kedua file ini ada di dalam package yang sama (service)
}

// constructor
func NewPurchaseReturnService(db *gorm.DB, purchaseReturnRepo repository.PurchaseReturnRepository, purchaseReturnItemRepo repository.PurchaseReturnItemRepository, itemVariantRepo repository.ItemVariantRepository, stockMovementRepo repository.StockMovementRepository, activityLogService ActivityLogService) PurchaseReturnService {
	return &purchaseReturnService{
		db:                     db,
		purchaseReturnRepo:     purchaseReturnRepo,
		purchaseReturnItemRepo: purchaseReturnItemRepo,
		itemVariantRepo:        itemVariantRepo,
		stockMovementRepo:      stockMovementRepo,
		activityLogService:     activityLogService,
	}
}

// struct method
func (s *purchaseReturnService) GetAllPurchaseReturns(ctx context.Context, query dto.GetAllPurchaseReturnsQuery) ([]dto.PurchaseReturnResponse, int64, error) {
	// get tenantID from context
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	purchaseReturn, total, err := s.purchaseReturnRepo.GetAllPurchaseReturns(ctx, tenantID, query)

	if err != nil {
		return nil, 0, err
	}

	// convert model to dto
	purchaseReturnDTO := helper.ConvertToDTOPurchaseReturnPlural(purchaseReturn)

	// jika semua sukses
	return purchaseReturnDTO, total, nil
}

func (s *purchaseReturnService) CreatePurchaseReturn(ctx context.Context, req dto.CreatePurchaseReturnRequest) (dto.PurchaseReturnResponse, error) {
	// begin transaction
	tx := s.db.Begin()

	// cek apakah ada error saat begin transaction
	if tx.Error != nil {
		return dto.PurchaseReturnResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	// ambil tenantID dan userID dari context
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	userID := ctx.Value(constants.ContextUserID).(string)

	// generate purchase return number
	returnNumber := fmt.Sprintf(
		"P-RETUR-%d",
		time.Now().Unix(),
	)

	// ========================================
	// CREATE PURCHASE RETURN HEADER -> master
	// ========================================
	purchaseReturn := model.PurchaseReturn{
		ID:           uuid.NewString(),
		TenantID:     tenantID,
		PurchaseID:   req.PurchaseID,
		ReturnNumber: returnNumber,
		Reason:       req.Reason,
		Notes:        req.Notes,
		CreatedBy:    userID,
	}

	// create header / master purchase return
	err := s.purchaseReturnRepo.CreatePurchaseReturn(ctx, tx, &purchaseReturn)
	if err != nil {
		tx.Rollback() // rollback jika operasi
		return dto.PurchaseReturnResponse{}, err
	}

	// loop purchase return items
	for _, item := range req.Items { // loop ke struct req.Items yang bersifat plural
		// ambil item variant dari db, harus trusted, jangan percaya input dari frontend
		itemVariant, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, item.ItemVariantID)
		if err != nil {
			return dto.PurchaseReturnResponse{}, err
		}

		purchaseReturnItems := model.PurchaseReturnItem{
			ID:               uuid.NewString(),
			TenantID:         tenantID,
			PurchaseReturnID: purchaseReturn.ID,
			ItemVariantID:    itemVariant.ID,
			Qty:              item.Qty,
			Notes:            item.Notes,
		}

		// simpan purchase return items ke db
		err = s.purchaseReturnItemRepo.CreatePurchaseReturnItem(ctx, tx, &purchaseReturnItems)

		if err != nil {
			tx.Rollback()
			return dto.PurchaseReturnResponse{}, err
		}

		// ========================================
		// CREATE STOCK MOVEMENT
		// stok keluar (retur), qty selalu negatif
		// ========================================

		movement := model.StockMovement{
			ID:            uuid.NewString(),
			TenantID:      tenantID,
			ItemVariantID: itemVariant.ID,
			MovementType:  constants.MovementPurchaseReturn,
			Qty:           -item.Qty, // selalu negatif karena retur mengurangi stok yang ada
			ReferenceType: "PURCHASE RETURN",
			ReferenceID:   purchaseReturn.ID,
			Notes:         "purchase return transaction",
			CreatedBy:     userID,
		}

		// simpan data ke tabel movement
		err = s.stockMovementRepo.CreateMovement(ctx, tx, &movement)
		if err != nil {
			tx.Rollback()
			return dto.PurchaseReturnResponse{}, err
		}
	}

	// commit transaction -> WAJIB
	err = tx.Commit().Error

	// di blok ini tidak perlu di rollback lagi, tx sudah di commit
	if err != nil {
		return dto.PurchaseReturnResponse{}, err
	}

	// setelah selesai commit transaction, eksekusi service CreateActivityLog()
	_ = s.activityLogService.CreateActivityLog( // ignore error
		ctx,
		"PURCHASE RETURN",
		"CREATE",
		fmt.Sprintf("Create Purchase Return %s", purchaseReturn.ReturnNumber),
		purchaseReturn.ID,           // id uuid
		purchaseReturn.ReturnNumber, // yang mudah dipahami manusia misalnya P-RETUR-1781140525
	)

	// get data by ID + preload relasi
	newPurchaseReturn, err := s.purchaseReturnRepo.GetPurchaseReturnByID(ctx, tenantID, purchaseReturn.ID)

	if err != nil {
		return dto.PurchaseReturnResponse{}, err
	}

	// convert model to dto
	purchaseReturnDTO := helper.ConvertToDTOPurchaseReturnSingle(newPurchaseReturn)

	return purchaseReturnDTO, nil
}

func (s *purchaseReturnService) GetPurchaseReturnByID(ctx context.Context, id string) (dto.PurchaseReturnResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	purchaseReturn, err := s.purchaseReturnRepo.GetPurchaseReturnByID(ctx, tenantID, id)
	if err != nil {
		return dto.PurchaseReturnResponse{}, err
	}

	// convert model to dto
	purchaseReturnDTO := helper.ConvertToDTOPurchaseReturnSingle(purchaseReturn)

	return purchaseReturnDTO, nil
}
