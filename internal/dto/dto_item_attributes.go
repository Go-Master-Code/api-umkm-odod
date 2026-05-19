package dto

// response json
type ItemAttributeResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
}

// create request
type CreateItemAttributeRequest struct {
	// ID       string `json:"id" binding:"required,uuid"` uuid dibuat backend, bukan request body
	Name string `json:"name" binding:"required,min=3,max=100"`
	// TenantID string `json:"tenant_id" binding:"required,uuid"` berasal dari jwt
}

// update request
type UpdateItemAttributeRequest struct {
	Name *string `json:"name" binding:"omitempty,min=3,max=100"`
	// TenantID *string `json:"tenant_id" binding:"omitempty,uuid"` jangan boleh update, patokan TenantID adalah jwt
}
