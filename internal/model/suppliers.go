package model

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID        string         `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID  string         `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant    Tenant         `json:"-" gorm:"foreignKey:TenantID"`
	Name      string         `json:"name" gorm:"type:varchar(150);not null"`
	Phone     string         `json:"phone" gorm:"type:varchar(30);not null"`
	Address   string         `json:"address" gorm:"type:text;not null"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
