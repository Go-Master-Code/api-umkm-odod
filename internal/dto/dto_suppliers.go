package dto

// response
type SupplierResponse struct {
	ID       string `json:"id"`
	TenantID string `json:"tenant_id"`
	Name     string `json:"tenant_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}

// create request
type CreateSupplierRequest struct {
	Name    string `json:"name" binding:"required,min=3,max=150"`
	Phone   string `json:"phone" binding:"required,min=8,max=30"`
	Address string `json:"address" binding:"required,max=300"`
}

// update request
type UpdateSupplierRequest struct {
	Name    *string `json:"name" binding:"omitempty,min=3,max=150"`
	Phone   *string `json:"phone" binding:"omitempty,min=8,max=30"`
	Address *string `json:"address" binding:"omitempty,max=300"`
}
