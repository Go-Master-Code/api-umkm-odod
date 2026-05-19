package dto

// response
type ItemVariantResponse struct {
	ID          string  `json:"id"`
	TenantID    string  `json:"tenant_id"`
	TenantName  string  `json:"tenant_name"`
	ItemID      string  `json:"item_id"`
	ItemName    string  `json:"item_name"`
	SKU         string  `json:"sku"`
	Barcode     string  `json:"barcode"`
	VariantName string  `json:"variant_name"`
	CostPrice   float64 `json:"cost_price"`
	IsActive    bool    `json:"is_active"`
}

// create request
type CreateItemVariantRequest struct {
	// ID          string  `json:"id" binding:"required,uuid"` tidak usah karena uuid di generate backend, bukan dari request body
	// TenantID    string  `json:"tenant_id" binding:"required,uuid"` jangan dari create request, berasal dari jwt harusnya
	ItemID      string  `json:"item_id" binding:"required,uuid"`
	SKU         string  `json:"sku" binding:"required,min=3,max=100"`
	Barcode     string  `json:"barcode" binding:"omitempty,max=100"` // tidak required
	VariantName string  `json:"variant_name" binding:"required,min=3,max=150"`
	CostPrice   float64 `json:"cost_price" binding:"required,gte=0"` // gte = nilai min = 0
	IsActive    bool    `json:"is_active"`
}

// update request
type UpdateItemVariantRequest struct {
	ItemID      *string  `json:"item_id" binding:"omitempty,uuid"`
	SKU         *string  `json:"sku" binding:"omitempty,min=3,max=100"`
	Barcode     *string  `json:"barcode" binding:"omitempty,max=100"`
	VariantName *string  `json:"variant_name" binding:"omitempty,min=3,max=150"`
	CostPrice   *float64 `json:"cost_price" binding:"omitempty,gte=0"`
	IsActive    *bool    `json:"is_active"` // update request var bool tidak perlu omitempty
}
