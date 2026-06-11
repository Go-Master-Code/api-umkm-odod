package repository

import (
	"context"
	"time"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type DashboardRepository interface {
	GetTodaySales(ctx context.Context, tenantID string) (float64, int64, error)
	GetTodayPurchases(ctx context.Context, tenantID string) (float64, int64, error)
	// GetLowStockCount(ctx context.Context, tenantID string) (int64, error)
	// GetTotalItems(ctx context.Context, tenantID string) (int64, error)
	// GetTotalVariants(ctx context.Context, tenantID string) (int64, error)
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
