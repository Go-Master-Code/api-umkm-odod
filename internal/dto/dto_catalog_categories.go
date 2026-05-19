package dto

// dto response
type CatalogCategoryResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
}

// create request
type CreateCatalogCategoryRequest struct {
	// ID       string `json:"id" binding:"required,uuid"`
	Name string `json:"name" binding:"required,min=3,max=100"`
	// TenantID string `json:"tenant_id" binding:"required,uuid"`
}

// update request wajib pointer
type UpdateCatalogCategoryRequest struct {
	Name *string `json:"name" binding:"omitempty,min=3,max=100"`
	// TenantID *string `json:"tenant_id" binding:"omitempty,uuid"`
}
