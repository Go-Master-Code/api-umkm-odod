package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ReportRepository interface {
	GetSalesReport(ctx context.Context, tenantID string, startDate string, endDate string) ([]model.Sale, error)
	GetSalesReportSummary(ctx context.Context, tenantID string, startDate string, endDate string) (*dto.SalesReportSummary, error)
	GetPurchaseReport(ctx context.Context, tenantID string, startDate string, endDate string) ([]model.Purchase, error)
	GetPurchaseReportSummary(ctx context.Context, tenantID string, startDate string, endDate string) (*dto.PurchaseReportSummary, error)
	GetStockReport(ctx context.Context, tenantID string, query dto.StockReportQuery) ([]dto.StockReportResponse, int64, error)
	GetStockReportSummary(ctx context.Context, tenantID string) (*dto.StockReportSummary, error)
}

// struct implementasi
type reportRepository struct {
	db *gorm.DB
}

// constructor
func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{
		db: db,
	}
}

func (r *reportRepository) GetSalesReport(ctx context.Context, tenantID string, startDate string, endDate string) ([]model.Sale, error) {
	var sales []model.Sale
	err := r.db.WithContext(ctx).Model(&model.Sale{}).
		Preload("Tenant").
		Preload("Cashier").
		Preload("SaleItems").
		Preload("SaleItems.Tenant").      // preload nested relation dari sale item
		Preload("SaleItems.ItemVariant"). // preload nested relation dari sale item
		Where("tenant_id = ? and DATE(created_at) BETWEEN ? AND ?", tenantID, startDate, endDate).
		Order("created_at DESC").
		Find(&sales).Error

	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *reportRepository) GetSalesReportSummary(ctx context.Context, tenantID string, startDate string, endDate string) (*dto.SalesReportSummary, error) {
	var salesReportSummary dto.SalesReportSummary
	err := r.db.WithContext(ctx).Model(&model.Sale{}).
		Where("tenant_id = ? AND DATE(created_at) BETWEEN ? AND ?", tenantID, startDate, endDate).
		Select(`count(*) AS total_transaction,
			COALESCE(SUM(subtotal),0) AS total_sales,
			COALESCE(SUM(discount_amount),0) AS total_discount,
			COALESCE(SUM(tax_amount),0) AS total_tax,
			COALESCE(SUM(grand_total),0) AS grand_total
		`).
		Scan(&salesReportSummary).Error

	if err != nil {
		return &dto.SalesReportSummary{}, err
	}

	return &salesReportSummary, nil
}

func (r *reportRepository) GetPurchaseReport(ctx context.Context, tenantID string, startDate string, endDate string) ([]model.Purchase, error) {
	var purchase []model.Purchase

	err := r.db.WithContext(ctx).Model(&model.Purchase{}).
		Preload("Tenant").
		Preload("Supplier").
		Preload("Creator").
		Preload("PurchaseItems").
		Preload("PurchaseItems.Tenant").      // preload nested relation dari purchase item
		Preload("PurchaseItems.ItemVariant"). // preload nested relation dari purchase item
		Where("tenant_id = ? AND DATE(created_at) BETWEEN ? AND ?", tenantID, startDate, endDate).
		Order("created_at DESC").
		Find(&purchase).Error

	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (r *reportRepository) GetPurchaseReportSummary(ctx context.Context, tenantID string, startDate string, endDate string) (*dto.PurchaseReportSummary, error) {
	var purchaseReportSummary dto.PurchaseReportSummary

	err := r.db.WithContext(ctx).Model(&model.Purchase{}).
		Where("tenant_id = ? AND DATE(created_at) BETWEEN ? AND ?", tenantID, startDate, endDate).
		Select(`count(*) AS total_transaction,
			COALESCE(SUM(subtotal),0) AS total_purchase,
			COALESCE(SUM(discount_amount),0) AS total_discount,
			COALESCE(SUM(tax_amount),0) AS total_tax,
			COALESCE(SUM(grand_total),0) AS grand_total
		`).
		Scan(&purchaseReportSummary).Error

	if err != nil {
		return &dto.PurchaseReportSummary{}, err
	}

	return &purchaseReportSummary, nil
}

func (r *reportRepository) GetStockReport(ctx context.Context, tenantID string, query dto.StockReportQuery) ([]dto.StockReportResponse, int64, error) {
	var result []dto.StockReportResponse
	var total int64

	offset := (query.Page - 1) * query.Limit

	baseQuery := r.db.WithContext(ctx).Model(&model.ItemVariant{}).Where("item_variants.tenant_id = ?", tenantID)

	// jika query search tidak kosong
	if query.Search != "" {
		search := "%" + query.Search + "%"
		baseQuery = baseQuery.Joins("LEFT JOIN catalog_items ON catalog_items.id = item_variants.item_id").
			Where("catalog_items.name LIKE ? OR item_variants.sku LIKE ?", search, search)

	}

	// count hasil query
	err := baseQuery.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// query untuk select data
	err = baseQuery.Select(`item_variants.id AS item_variant_id,
			catalog_items.name AS item_name,
			item_variants.variant_name AS variant_name,
			item_variants.sku AS sku,
			item_variants.minimum_stock AS minimum_stock,
			COALESCE(sum(stock_movements.qty),0) AS current_stock`).
		Joins("LEFT JOIN catalog_items ON catalog_items.id = item_variants.item_id").
		Joins("LEFT JOIN stock_movements ON stock_movements.item_variant_id = item_variants.id").
		Group(`item_variants.id, catalog_items.name, item_variants.variant_name, item_variants.sku, item_variants.minimum_stock`). // untuk klausa GROUP BY harus pakai nama kolom ASLI bukan AS atau ALIAS
		Order("catalog_items.name").
		Limit(query.Limit).
		Offset(offset).
		Scan(&result).Error

	if err != nil {
		return nil, 0, err
	}

	// jika sukses
	return result, total, nil
}

func (r *reportRepository) GetStockReportSummary(ctx context.Context, tenantID string) (*dto.StockReportSummary, error) {
	var summary dto.StockReportSummary

	// cari total variant
	err := r.db.WithContext(ctx).Model(&model.ItemVariant{}).Where("tenant_id = ?", tenantID).Count(&summary.TotalVariants).Error
	if err != nil {
		return nil, err
	}

	// low stock items
	err = r.db.WithContext(ctx).Model(&model.ItemVariant{}).
		Joins("LEFT JOIN stock_movements on stock_movements.item_variant_id = item_variants.id").
		Where("item_variants.tenant_id = ?", tenantID).
		Group("item_variants.id").
		Having("COALESCE(sum(stock_movements.qty),0) <= item_variants.minimum_stock").
		Count(&summary.LowStockItems).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil
}
