package model

import "time"

type PurchaseReturnItem struct {
	ID               string         `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID         string         `json:"tenant_id" gorm:"type:char(36);not null;index"`
	PurchaseReturnID string         `json:"purchase_return_id" gorm:"type:char(36);not null"`
	PurchaseReturn   PurchaseReturn `json:"-" gorm:"foreignKey:PurchaseReturnID"` // reverse relation ke model PurchaseReturn
	ItemVariantID    string         `json:"item_variant_id" gorm:"type:char(36);not null"`
	ItemVariant      ItemVariant    `json:"-" gorm:"foreignKey:ItemVariantID;not null"`
	Qty              float64        `json:"qty" gorm:"type:decimal(18,2);not null"`
	Notes            string         `json:"notes" gorm:"type:text"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
}

func (PurchaseReturnItem) TableName() string {
	return "purchase_return_items"
}
