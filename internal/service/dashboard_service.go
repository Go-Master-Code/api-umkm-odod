package service

import (
	"context"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"
)

// interface
type DashboardService interface {
	GetSummary(ctx context.Context) (dto.DashBoardSummaryResponse, error)
	GetDailySalesChart(ctx context.Context) ([]dto.DailySalesChartResponse, error)
	GetDailyPurchaseChart(ctx context.Context) ([]dto.DailyPurchaseChartResponse, error)
}

// struct implementasi
type dashboardService struct {
	dashboardRepo     repository.DashboardRepository
	stockMovementRepo repository.StockMovementRepository
	itemVariantRepo   repository.ItemVariantRepository
	catalogItemRepo   repository.CatalogItemRepository
	supplierRepo      repository.SupplierRepository
}

// constructor
func NewDashboardService(dashboardRepo repository.DashboardRepository, stockMovementRepo repository.StockMovementRepository, itemVariantRepo repository.ItemVariantRepository, catalogItemRepo repository.CatalogItemRepository, supplierRepo repository.SupplierRepository) DashboardService {
	return &dashboardService{
		dashboardRepo:     dashboardRepo,
		stockMovementRepo: stockMovementRepo,
		itemVariantRepo:   itemVariantRepo,
		catalogItemRepo:   catalogItemRepo,
		supplierRepo:      supplierRepo,
	}
}

// struct method
func (s *dashboardService) GetSummary(ctx context.Context) (dto.DashBoardSummaryResponse, error) {
	// get tenantID by context
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// =========================
	// ==========SALES==========
	// =========================
	totalSales, totalTransactions, err := s.dashboardRepo.GetTodaySales(ctx, tenantID)
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// =============================
	// ==========LOW STOCK==========
	// =============================
	variants, err := s.itemVariantRepo.GetAllItemVariants(ctx, tenantID)
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	var lowStockCount int64

	for _, variant := range variants {
		// ambil stok realtime
		// kalau tx ada -> pakai transaction
		// kalau tx nil -> pakai repository db biasa, dan query akan menggunakan r.db, tidak pakai mode transaction (tx)
		currentStock, err := s.stockMovementRepo.GetCurrentStock(ctx, tenantID, nil, variant.ID)

		if err != nil {
			return dto.DashBoardSummaryResponse{}, err
		}

		if currentStock < variant.MinimumStock {
			lowStockCount += 1 // increment jumlah barang kategori low stock
		}
	}

	// ============================
	// ==========PURCHASE==========
	// ============================
	totalPurchases, totalPurchaseTransactions, err := s.dashboardRepo.GetTodayPurchases(ctx, tenantID)
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// =============================
	// ==========SUPPLIERS==========
	// =============================
	suppliers, err := s.supplierRepo.GetSuppliers(ctx, tenantID, "")
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// convert int ke int64
	totalSuppliers := int64(len(suppliers))

	// ==============================================
	// ==========CATALOG ITEM (JUMLAH ITEM)==========
	// ==============================================
	catalogItems, err := s.catalogItemRepo.GetCatalogItems(ctx, tenantID, "")
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// convert int ke int64
	totalItems := int64(len(catalogItems))

	// =================================
	// ==========ITEM VARIANTS==========
	// =================================
	itemVariants, err := s.itemVariantRepo.GetItemVariants(ctx, tenantID, "")
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// convert int ke int64
	TotalVariants := int64(len(itemVariants))

	// mapping hasil ke dto
	dashboardSummary := dto.DashBoardSummaryResponse{
		TodaySales:               totalSales,
		TodayTransactions:        totalTransactions,
		LowStockCount:            lowStockCount,
		TodayPurchase:            totalPurchases,
		TodayPurchaseTransaction: totalPurchaseTransactions,
		TotalSuppliers:           totalSuppliers,
		TotalItems:               totalItems,
		TotalVariants:            TotalVariants,
	}

	return dashboardSummary, nil
}

func (s *dashboardService) GetDailySalesChart(ctx context.Context) ([]dto.DailySalesChartResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	dailySalesChart, err := s.dashboardRepo.GetDailySalesChart(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return dailySalesChart, nil
}

func (s *dashboardService) GetDailyPurchaseChart(ctx context.Context) ([]dto.DailyPurchaseChartResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	dailyPurchaseChart, err := s.dashboardRepo.GetDailyPurchaseChart(ctx, tenantID)

	if err != nil {
		return nil, err
	}

	return dailyPurchaseChart, nil
}
