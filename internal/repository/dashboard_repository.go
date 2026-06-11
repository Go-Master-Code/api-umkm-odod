package repository

import (
	"context"
	"time"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type DashboardRepository interface {
	GetTodaySales(ctx context.Context, tenantID string) (float64, int64, error)
	GetTodayPurchases(ctx context.Context, tenantID string) (float64, int64, error)
	GetDailySalesChart(ctx context.Context, tenantID string) ([]dto.DailySalesChartResponse, error)
	GetDailyPurchaseChart(ctx context.Context, tenantID string) ([]dto.DailyPurchaseChartResponse, error)
	GetTopSellingProducts(ctx context.Context, tenantID string) ([]dto.TopSellingProductsResponse, error)
}

// struct implementasi
type dashboardRepository struct {
	db *gorm.DB
}

// constructor
func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{
		db: db,
	}
}

// struct method
func (r *dashboardRepository) GetTodaySales(ctx context.Context, tenantID string) (float64, int64, error) {
	var totalSales float64
	var totalTransactions int64

	today := time.Now().Format("2006-01-02")

	// jumlahkan semua grand total untuk mendapatkan totalSales
	err := r.db.WithContext(ctx).Model(&model.Sale{}).
		Where("tenant_id = ? and DATE(created_at) = ?", tenantID, today).
		Select("COALESCE(SUM(grand_total),0)").
		Scan(&totalSales).Error

	if err != nil {
		return 0, 0, err
	}

	// query untuk mencari jumlah transaksi sales yang terjadi
	err = r.db.WithContext(ctx).Model(&model.Sale{}).
		Where("tenant_id = ? and DATE(created_at) = ?", tenantID, today).
		Count(&totalTransactions).Error

	if err != nil {
		return 0, 0, err
	}

	// jika kedua query sukses
	return totalSales, totalTransactions, nil
}

func (r *dashboardRepository) GetTodayPurchases(ctx context.Context, tenantID string) (float64, int64, error) {
	var totalPurchases float64
	var totalPurchaseTransactions int64

	today := time.Now().Format("2006-01-02")

	// query total purchase today
	err := r.db.WithContext(ctx).Model(&model.Purchase{}).
		Where("tenant_id = ? AND DATE(created_at) = ?", tenantID, today).
		Select("COALESCE(SUM(grand_total),0)").Scan(&totalPurchases).Error

	if err != nil {
		return 0, 0, err
	}

	// query untuk mengambil jumlah purchase transaction hari ini
	err = r.db.WithContext(ctx).Model(&model.Purchase{}).
		Where("tenant_id = ? and DATE(created_at) = ?", tenantID, today).
		Count(&totalPurchaseTransactions).Error

	if err != nil {
		return 0, 0, err
	}

	return totalPurchases, totalPurchaseTransactions, nil
}

func (r *dashboardRepository) GetDailySalesChart(ctx context.Context, tenantID string) ([]dto.DailySalesChartResponse, error) {
	var sales []dto.DailySalesChartResponse

	err := r.db.WithContext(ctx).Model(&model.Sale{}).
		Where("tenant_id = ? AND created_at > DATE_SUB(NOW(), INTERVAL 7 DAY)", tenantID).
		Select("DATE(created_at) AS date, SUM(grand_total) AS total_sales").
		Group("DATE(created_at)").
		Order("DATE(created_at)").
		Scan(&sales).Error // pakai scan jangan find, karena di query ini pakai model &Model.Sale{} Karena Anda tidak sedang mengambil entity Sale, melainkan hasil agregasi custom.

	// field di group dan order harus sesuai dengan select, tidak boleh created_at saja

	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *dashboardRepository) GetDailyPurchaseChart(ctx context.Context, tenantID string) ([]dto.DailyPurchaseChartResponse, error) {
	var purchase []dto.DailyPurchaseChartResponse

	err := r.db.WithContext(ctx).Model(&model.Purchase{}).
		Where("tenant_id = ? AND created_at > DATE_SUB(NOW(), INTERVAL 7 DAY)", tenantID).
		Select("DATE(created_at) AS date, SUM(grand_total) AS total_purchase").
		Group("DATE(created_at)").
		Order("DATE(created_at)").
		Scan(&purchase).Error // pakai scan jangan find, karena di query ini pakai model &Model.Purchase{} Karena Anda tidak sedang mengambil entity Purchase, melainkan hasil agregasi custom.

	// field di group dan order harus sesuai dengan select, tidak boleh created_at saja

	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (r *dashboardRepository) GetTopSellingProducts(ctx context.Context, tenantID string) ([]dto.TopSellingProductsResponse, error) {
	var topSellingProducts []dto.TopSellingProductsResponse

	// nama field di query sql pakai AS agar sesuai dengan nama field json pada topSellingProducts
	err := r.db.WithContext(ctx).
		Model(model.SaleItem{}).
		Where("tenant_id = ?", tenantID).
		Select(`item_variant_id, item_name_snapshot AS item_name, variant_name_snapshot AS variant_name, SUM(qty) AS qty_sold`).
		Group(`item_variant_id, item_name_snapshot, variant_name_snapshot`).
		Order("qty_sold desc"). // urutkan dari yang terjual terbanyak
		Limit(5).               // batasi hanya 5 produk
		Scan(&topSellingProducts).Error

	if err != nil {
		return []dto.TopSellingProductsResponse{}, err
	}

	return topSellingProducts, nil
}
