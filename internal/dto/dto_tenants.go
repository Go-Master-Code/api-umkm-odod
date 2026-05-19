package dto

// dto response
type TenantResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// create request
type CreateTenantRequest struct {
	// ID      string `json:"id" binding:"required,uuid"` dibuat di backend, bukan request body
	Name    string `json:"name" binding:"required,min=3,max=150"`
	Phone   string `json:"phone" binding:"required,min=8,max=30"`
	Address string `json:"address" binding:"required,max=500"`
}

// update request wajib pointer
type UpdateTenantRequest struct {
	Name    *string `json:"name" binding:"omitempty,min=3,max=150"`
	Phone   *string `json:"phone" binding:"omitempty,min=8,max=30"`
	Address *string `json:"address" binding:"omitempty,max=500"`
}
