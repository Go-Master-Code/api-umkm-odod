package model

import (
	"time"

	"gorm.io/gorm"
)

type CatalogItem struct {
	ID              string          `json:"id" gorm:"type:char(36);primaryKey"`            // format jika pakai uuid
	TenantID        string          `json:"tenant_id" gorm:"type:char(36);not null;index"` // untuk optimalisasi query karena nanti akan pakai where tenant_id = ?
	Tenant          Tenant          `json:"-" gorm:"foreignKey:TenantID"`
	CategoryID      string          `json:"category_id" gorm:"type:char(36);not null;index"`
	CatalogCategory CatalogCategory `json:"-" gorm:"foreignKey:CategoryID"`
	Name            string          `json:"name" gorm:"type:varchar(200);not null"`
	Description     string          `json:"description"`
	IsActive        bool            `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt       gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (CatalogItem) TableName() string {
	return "catalog_items"
}
