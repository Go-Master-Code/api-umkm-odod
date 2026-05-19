package model

type ItemVariantAttributeValue struct {
	TenantID         string             `json:"tenant_id" gorm:"type:char(36);not null;primaryKey"`
	Tenant           Tenant             `json:"-" gorm:"foreignKey:TenantID"`
	VariantID        string             `json:"variant_id" gorm:"type:char(36);primaryKey"`
	Variant          ItemVariant        `json:"-" gorm:"foreignKey:VariantID"`
	AttributeValueID string             `json:"attribute_value_id" gorm:"type:char(36);primaryKey"`
	AttributeValue   ItemAttributeValue `json:"-" gorm:"foreignKey:AttributeValueID"`
}

func (ItemVariantAttributeValue) TableName() string {
	return "item_variant_attribute_values"
}
