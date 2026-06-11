package dto

type SaleReportQuery struct {
	StartDate string `form:"start_date"` // untuk query parameter URL, bind c.json.shouldbindquery, harus pakai tag form, bukan json
	EndDate   string `form:"end_date"`
}

// summary
type SalesReportSummary struct {
	TotalTransaction int64   `json:"total_transaction"`
	TotalSales       float64 `json:"total_sales"`
	TotalDiscount    float64 `json:"total_discount"`
	TotalTax         float64 `json:"total_tax"`
	GrandTotal       float64 `json:"grand_total"`
}

type SalesReportResponse struct {
	Summary      SalesReportSummary `json:"summary"`
	Transactions []SaleResponse     `json:"transactions"` // master sales berikut detail (field Items di dto SaleResponse)
}
