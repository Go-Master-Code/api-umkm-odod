package dto

import "time"

// response
type StockMovementResponse struct {
	ID              string    `json:"id"`
	TenantID        string    `json:"tenant_id"`
	TenantName      string    `json:"tenant_name"`
	ItemVariantID   string    `json:"item_variant_id"`
	ItemVariantName string    `json:"item_variant_name"`
	MovementType    string    `json:"movement_type"`
	Qty             float64   `json:"qty"`
	ReferenceType   string    `json:"reference_type"`
	ReferenceID     string    `json:"reference_id"`
	Notes           string    `json:"notes"`
	CreatedBy       string    `json:"created_by"`
	CreatedByName   string    `json:"created_by_name"`
	CreatedAt       time.Time `json:"created_at"`
}

// kode untuk create request dibuat di internal/payloads/stock_movement_payload.go

// kode request dto untuk add stock, reduce stock, dan cek stok dibuat di file ini

// add stock -> dipanggil di handler sebagai req (param untuk service)
type AddStockRequest struct {
	ItemVariantID string  `json:"item_variant_id" binding:"required,uuid"`
	Qty           float64 `json:"qty" binding:"required,gt=0"`
	Notes         string  `json:"notes" binding:"omitempty,max=200"`
}

// reduce stock
type ReduceStockRequest struct {
	ItemVariantID string  `json:"item_variant_id" binding:"required,uuid"`
	Qty           float64 `json:"qty" binding:"required,gt=0"`
	Notes         string  `json:"notes" binding:"omitempty,max=200"`
}

// current stock response
type CurrentStockResponse struct {
	ItemVariantID string  `json:"item_variant_id"`
	CurrentStock  float64 `json:"current_stock"`
}
