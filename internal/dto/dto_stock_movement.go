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

// kode untuk create request dibuat di internal/payloads/stock_movement_payload.dt
