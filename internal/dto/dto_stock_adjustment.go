package dto

import "time"

/*
	Stock adjustment diperuntukan apabila terjadi beberapa kemungkinan seperti :
	- barang hilang
	- barang rusak
	- barang kadaluarsa
	- koreksi stok (hasil stock opname)
	- salah input
*/

type CreateStockAdjustmentRequest struct {
	ItemVariantID string  `json:"item_variant_id" binding:"required,uuid"`
	Qty           float64 `json:"qty" binding:"required,gt=0"`
	Reason        string  `json:"reason" binding:"required,max=300"`
	Type          string  `json:"type" binding:"required,oneof=ADD REDUCE"`
	Notes         string  `json:"notes" binding:"omitempty,max=300"`
}

type StockAdjustmenResponse struct {
	ID              string    `json:"id"`
	ItemVariantID   string    `json:"item_variant_id"`
	ItemVariantName string    `json:"item_variant_name"`
	Type            string    `json:"type"` // add / reduce
	Qty             float64   `json:"qty"`
	Reason          string    `json:"reason"`
	Notes           string    `json:"notes"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
}
