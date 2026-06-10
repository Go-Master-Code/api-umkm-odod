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
	CreatePurchaseReturn(ctx context.Context, req dto.CreatePurchaseReturnRequest) (dto.PurchaseReturnResponse, error)
}

// struct implementasi
type purchaseReturnService struct {
	db                     *gorm.DB
	repoPurchaseReturn     repository.PurchaseReturnRepository
	repoPurchaseReturnItem repository.PurchaseReturnItemRepository
	repoStockMovement      repository.StockMovementRepository
}

// constructor
func NewPurchaseReturnService(db *gorm.DB, repoPurchaseReturn repository.PurchaseReturnRepository, repoPurchaseReturnItem repository.PurchaseReturnItemRepository, repoStockMovement repository.StockMovementRepository) PurchaseReturnService {
	return &purchaseReturnService{
		db:                     db,
		repoPurchaseReturn:     repoPurchaseReturn,
		repoPurchaseReturnItem: repoPurchaseReturnItem,
		repoStockMovement:      repoStockMovement,
	}
}

// struct method
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

	err := s.repoPurchaseReturn.CreatePurchaseReturn(ctx, tx, &purchaseReturn)
	if err != nil {
		tx.Rollback() // rollback jika operasi
		return dto.PurchaseReturnResponse{}, err
	}

	// get data by ID + preload relasi
	newPurchaseReturn, err := s.repoPurchaseReturn.GetPurchaseReturnByID(ctx, tenantID, purchaseReturn.ID)

	if err != nil {
		return dto.PurchaseReturnResponse{}, err
	}

	// convert model to dto
	purchaseReturnDTO := helper.ConvertToDTOPurchaseReturnSingle(newPurchaseReturn)
	return purchaseReturnDTO, nil
}
