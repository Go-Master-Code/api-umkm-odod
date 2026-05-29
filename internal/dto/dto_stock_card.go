package dto

import "time"

// response per row kartu stok
type StockCardResponse struct {
	MovementDate  time.Time `json:"movement_date"`
	MovementType  string    `json:"movement_type"`
	QtyIn         float64   `json:"qty_in"`
	QtyOut        float64   `json:"qty_out"`
	Balance       float64   `json:"balance"`        // saldo setelah movement barang
	ReferenceType string    `json:"reference_type"` // SALE / PURCHASE / ADJUSTMENT
	ReferenceID   string    `json:"reference_id"`
	Notes         string    `json:"notes"`
	CreatedByName string    `json:"created_by_name"`
}
