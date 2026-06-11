package dto

type DashBoardSummaryResponse struct {
	TodaySales               float64 `json:"today_sales"`
	TodayTransactions        int64   `json:"today_transactions"`
	TodayPurchase            float64 `json:"today_purchase"`
	TodayPurchaseTransaction int64   `json:"today_purchase_transaction"`
	LowStockCount            int64   `json:"low_stock_count"`
	TotalItems               int64   `json:"total_items"`
	TotalVariants            int64   `json:"total_variants"`
	TotalSuppliers           int64   `json:"total_suppliers"`
}

// untuk chart total sales
type DailySalesChartResponse struct {
	Date       string  `json:"date"`
	TotalSales float64 `json:"total_sales"`
}

// untuk chart total purchase
type DailyPurchaseChartResponse struct {
	Date          string  `json:"date"`
	TotalPurchase float64 `json:"total_purchase"`
}
