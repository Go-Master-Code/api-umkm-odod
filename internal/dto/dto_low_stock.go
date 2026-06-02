package dto

type LowStockResponse struct {
	ItemVariantID string  `json:"item_variant_id"`
	ItemName      string  `json:"item_name"`
	VariantName   string  `json:"variant_name"`
	SKU           string  `json:"sku"`
	CurrentStock  float64 `json:"current_stock"`
	MinimumStock  float64 `json:"minimum_stock"`
}
