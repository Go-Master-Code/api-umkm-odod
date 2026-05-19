package model

import "time"

type PriceHistory struct {
	ID            string      `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID      string      `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant        Tenant      `json:"-" gorm:"foreignKey:TenantID"`
	ItemVariantID string      `json:"item_variant_id" gorm:"type:char(36);not null;index"`
	ItemVariant   ItemVariant `json:"-" gorm:"foreignKey:ItemVariantID"`
	PriceType     string      `json:"price_type" gorm:"type:varchar(50);not null;index"`
	Price         float64     `json:"price" gorm:"type:decimal(18,2);not null"`
	EffectiveDate time.Time   `json:"effective_date" gorm:"not null;index"`
	CreatedBy     string      `json:"created_by" gorm:"type:char(36);not null;index"`
	CreatedByUser string      `json:"-" gorm:"foreignKey:CreatedBy"`
	CreatedAt     time.Time   `gorm:"column:created_at;autoCreateTime"`
}

func (PriceHistory) TableName() string {
	return "price_history"
}
