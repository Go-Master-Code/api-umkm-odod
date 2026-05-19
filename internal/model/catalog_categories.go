package model

import (
	"time"

	"gorm.io/gorm"
)

type CatalogCategory struct {
	ID        string         `json:"id" gorm:"type:char(36);primaryKey"` // format jika pakai uuid
	Name      string         `json:"name" gorm:"type:varchar(100);not null"`
	TenantID  string         `json:"tenant_id" gorm:"type:char(36);not null;"`
	Tenant    Tenant         `json:"-" gorm:"foreignKey:TenantID"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (CatalogCategory) TableName() string {
	return "catalog_categories"
}
