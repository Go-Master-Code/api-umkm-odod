package dto

// response
type RoleResponse struct {
	ID         string `json:"id"`
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Name       string `json:"name"`
}

// create request
type CreateRoleRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}

// update request
type UpdateRoleRequest struct {
	Name *string `json:"name" binding:"omitempty,min=3,max=100"`
}
