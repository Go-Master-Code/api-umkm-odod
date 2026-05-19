package dto

type ItemVariantAttributeValueResponse struct {
	TenantID         string `json:"tenant_id"`
	TenantName       string `json:"tenant_name"`
	VariantID        string `json:"variant_id"`
	VariantName      string `json:"variant_name"`
	AttributeValueID string `json:"attribute_value_id"`
	AttributeValue   string `json:"attribute_value"`
}

type CreateItemVariantAttributeValue struct {
	TenantID         string `json:"tenant_id" binding:"required,uuid"`
	VariantID        string `json:"variant_id" binding:"required,uuid"`
	AttributeValueID string `json:"attribute_value_id" binding:"required,uuid"`
}

type UpdateItemVariantAttributeValue struct {
	TenantID         string `json:"tenant_id" binding:"omitempty,uuid"`
	VariantID        string `json:"variant_id" binding:"omitempty,uuid"`
	AttributeValueID string `json:"attribute_value_id" binding:"omitempty,uuid"`
}
