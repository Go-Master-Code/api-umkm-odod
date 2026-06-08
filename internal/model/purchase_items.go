package model

import "time"

type PurchaseItem struct {
	ID                  string      `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID            string      `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant              Tenant      `json:"-" gorm:"foreignKey:TenantID"`
	PurchaseID          string      `json:"purchase_id" gorm:"type:char(36);not null;index"`
	Purchase            Purchase    `json:"-" gorm:"foreignKey:PurchaseID"` // reverse relation to model purchase
	ItemVariantID       string      `json:"item_variant_id" gorm:"type:char(36);not null;index"`
	ItemVariant         ItemVariant `json:"-" gorm:"foreignKey:ItemVariantID"` // reverse relation to model ItemVariant
	ItemNameSnapshot    string      `json:"item_name_snapshot" gorm:"type:varchar(200);not null"`
	VariantNameSnapshot string      `json:"variant_name_snapshot" gorm:"type:varchar(150)"`
	SKUSnapshot         string      `json:"sku_snapshot" gorm:"type:varchar(100);not null"`
	Qty                 float64     `json:"qty" gorm:"type:decimal(18,2);not null"`
	CostPrice           float64     `json:"cost_price" gorm:"type:decimal(18,2);not null"`
	DiscountAmount      float64     `json:"discount_amount" gorm:"type:decimal(18,2);not null"`
	Subtotal            float64     `json:"subtotal" gorm:"type:decimal(18,2);not null"`
	CreatedAt           time.Time   `gorm:"column:created_at;autoCreateTime"`
}

func (PurchaseItem) TableName() string {
	return "purchase_items"
}
