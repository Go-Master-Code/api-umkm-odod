package dto

type CatalogItemResponse struct {
	ID           string `json:"id"`
	TenantID     string `json:"tenant_id"`
	TenantName   string `json:"tenant_name"`
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsActive     bool   `json:"is_active"`
}

// create request
type CreateCatalogItemRequest struct {
	// ID          string `json:"id" binding:"required,uuid"`
	// TenantID    string `json:"tenant_id" binding:"required"`
	CategoryID  string `json:"category_id" binding:"required,uuid"`
	Name        string `json:"name" binding:"required,min=3,max=200"`
	Description string `json:"description" binding:"max=200"`
	IsActive    bool   `json:"is_active"`
}

// update request
type UpdateCatalogItemRequest struct {
	// TenantID    *string `json:"tenant_id" binding:"omitempty"` tidak boleh update TenantID sembarangan
	CategoryID  *string `json:"category_id" binding:"omitempty,uuid"`
	Name        *string `json:"name" binding:"omitempty,min=3,max=200"`
	Description *string `json:"description" binding:"omitempty,max=200"`
	IsActive    *bool   `json:"is_active" binding:"omitempty"`
}
