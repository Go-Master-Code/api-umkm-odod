package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string         `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID    string         `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant      Tenant         `json:"-" gorm:"foreignKey:TenantID"`
	RoleID      string         `json:"role_id" gorm:"type:char(36);not null;index"`
	Role        Role           `json:"-" gorm:"foreignKey:RoleID"`
	FullName    string         `json:"full_name" gorm:"type:varchar(150);not null;index"`
	Username    string         `json:"username" gorm:"type:varchar(100);not null"`
	Password    string         `json:"-" gorm:"type:char(60);not null"`
	Phone       string         `json:"phone" gorm:"type:varchar(30);not null"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	LastLoginAt *time.Time     `json:"last_login_at"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
