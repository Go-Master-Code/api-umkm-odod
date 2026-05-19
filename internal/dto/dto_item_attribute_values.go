package dto

// response
type ItemAttributeValueResponse struct {
	ID            string `json:"id"`
	TenantID      string `json:"tenant_id"`
	TenantName    string `json:"tenant_name"`
	AttributeID   string `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
	Value         string `json:"value"`
}

// create request
type CreateItemAttributeValueRequest struct {
	// ID          string `json:"id" binding:"required,uuid"`
	// TenantID    string `json:"tenant_id" binding:"required,uuid"`
	AttributeID string `json:"attribute_id" binding:"required,uuid"`
	Value       string `json:"value" binding:"required,min=1,max=100"` // attribute value bisa pendek, misal S, M, L
}

// update request
type UpdateItemAttributeValueRequest struct {
	// TenantID    *string `json:"tenant_id" binding:"omitempty,uuid"`
	// AttributeID *string `json:"attribute_id" binding:"omitempty,uuid"` hati-hati attribute bisa berat, warna, rasa, jangan boleh update misal 250gr harusnya atribute berat, diganti jadi rasa
	Value *string `json:"value" binding:"omitempty,min=1,max=100"`
}
