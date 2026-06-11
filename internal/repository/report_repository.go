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
