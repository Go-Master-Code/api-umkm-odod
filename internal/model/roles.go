package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string         `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID  string         `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant    Tenant         `json:"-" gorm:"foreignKey:TenantID"`
	Name      string         `json:"name" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Role) TableName() string {
	return "roles"
}
