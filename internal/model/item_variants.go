package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemVariant struct {
	ID           string         `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID     string         `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant       Tenant         `json:"-" gorm:"foreignKey:TenantID"`
	ItemID       string         `json:"item_id" gorm:"type:char(36);not null;index"`
	Item         CatalogItem    `json:"-" gorm:"foreignKey:ItemID"` // nama var Item yang di preload di repo ItemVariant
	SKU          string         `json:"sku" gorm:"type:varchar(100);not null"`
	Barcode      string         `json:"barcode" gorm:"type:varchar(100);index"` // pakai index karena akan sering scanner barcode / pencarian kasir
	VariantName  string         `json:"variant_name" gorm:"type:varchar(150);not null"`
	CostPrice    float64        `json:"cost_price" gorm:"type:decimal(18,2);not null"`
	SellingPrice float64        `json:"selling_price" gorm:"type:decimal(18,2);not null;default:0"`
	MinimumStock float64        `json:"minimum_stock" gorm:"type:decimal(18,2);default:0"`
	IsActive     bool           `json:"is_active" gorm:"default:true"` // bool lebih baik diberi default value
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ItemVariant) TableName() string {
	return "item_variants"
}
