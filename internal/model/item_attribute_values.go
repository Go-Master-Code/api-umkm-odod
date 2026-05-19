package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemAttributeValue struct {
	ID            string         `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID      string         `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant        Tenant         `json:"-" gorm:"foreignKey:TenantID"`
	AttributeID   string         `json:"attribute_id" gorm:"type:char(36);not null;index"`
	ItemAttribute ItemAttribute  `json:"-" gorm:"foreignKey:AttributeID"`
	Value         string         `json:"value" gorm:"type:varchar(100);not null;index"` // memudahkan filter, search variant
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ItemAttributeValue) TableName() string {
	return "item_attribute_values"
}
