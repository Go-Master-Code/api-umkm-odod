package dto

// dto response
type TenantResponse struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	Address       string  `json:"address"`
	Email         string  `json:"email"`
	TaxPercentage float64 `json:"tax_percentage"`
	ReceiptFooter string  `json:"receipt_footer"`
}

// create request
type CreateTenantRequest struct {
	// ID      string `json:"id" binding:"required,uuid"` dibuat di backend, bukan request body
	Name          string  `json:"name" binding:"required,min=3,max=150"`
	Phone         string  `json:"phone" binding:"required,min=8,max=30"`
	Address       string  `json:"address" binding:"required,max=500"`
	Email         string  `json:"email" binding:"omitempty,max=150"`
	TaxPercentage float64 `json:"tax_percentage" binding:"required,gt=0"`
	ReceiptFooter string  `json:"receipt_footer" binding:"omitempty,max=300"`
}

// update request wajib pointer
type UpdateTenantRequest struct {
	Name          *string  `json:"name" binding:"omitempty,min=3,max=150"`
	Phone         *string  `json:"phone" binding:"omitempty,min=8,max=30"`
	Address       *string  `json:"address" binding:"omitempty,max=500"`
	Email         *string  `json:"email" binding:"omitempty,max=150"`
	TaxPercentage *float64 `json:"tax_percentage" binding:"omitempty,gt=0"`
	ReceiptFooter *string  `json:"receipt_footer" binding:"omitempty,max=300"`
}
