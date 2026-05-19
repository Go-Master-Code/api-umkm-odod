package model

import "time"

type StockMovement struct {
	ID            string      `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID      string      `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant        Tenant      `json:"-" gorm:"foreignKey:TenantID"`
	ItemVariantID string      `json:"item_variant_id" gorm:"type:char(36);not null;index"`
	ItemVariant   ItemVariant `json:"-" gorm:"foreignKey:ItemVariantID"`
	MovementType  string      `json:"movement_type" gorm:"type:varchar(50);not null;index"`
	Qty           float64     `json:"qty" gorm:"type:decimal(18,2);not null"`
	ReferenceType string      `json:"reference_type" gorm:"type:varchar(50)"`
	ReferenceID   string      `json:"reference_id" gorm:"type:char(36);index"` // karena akan sering dicari misal where reference_id=?
	Notes         string      `json:"notes" gorm:"type:varchar(200)"`          // tipe data mysql text boleh diganti pakai varchar(200) agar memori lebih efisien
	CreatedBy     string      `json:"created_by" gorm:"type:char(36);not null;index"`
	CreatedByUser User        `json:"-" gorm:"foreignKey:CreatedBy"`
	CreatedAt     time.Time   `gorm:"column:created_at;autoCreateTime"`
}

func (StockMovement) TableName() string {
	return "stock_movements"
}
