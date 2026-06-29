package model

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID                  string               `json:"id" gorm:"type:char(36);primaryKey"`
	Name                string               `json:"name" gorm:"type:varchar(150);not null"`
	Phone               string               `json:"phone" gorm:"type:varchar(30);not null"`
	Email               string               `json:"email" gorm:"type:varchar(150)"`
	Address             string               `json:"address" gorm:"type:text;not null"`
	TaxPercentage       float64              `json:"tax_percentage" gorm:"type:decimal(5,2);not null"`
	ReceiptFooter       string               `json:"receipt_footer" gorm:"type:text"`
	CatalogCategories   []CatalogCategory    `json:"-"` // reverse relation 1 to many (1 tenant bisa punya banyak catalog categories)
	CatalogItems        []CatalogItem        `json:"-"`
	ItemAttributes      []ItemAttribute      `json:"-"`
	ItemAttributeValues []ItemAttributeValue `json:"-"`
	ItemVariants        []ItemVariant        `json:"-"`
	CreatedAt           time.Time            `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time            `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt           gorm.DeletedAt       `gorm:"column:deleted_at"`
}

func (Tenant) TableName() string {
	return "tenants"
}
