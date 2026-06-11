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
