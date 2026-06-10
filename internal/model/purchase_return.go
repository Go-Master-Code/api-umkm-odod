package model

import "time"

type PurchaseReturn struct {
	ID                  string               `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID            string               `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant              Tenant               `json:"-" gorm:"foreignKey:TenantID"`
	PurchaseID          string               `json:"purchase_id" gorm:"type:char(36);not null;index"`
	Purchase            Purchase             `json:"-" gorm:"foreignKey:PurchaseID"`
	ReturnNumber        string               `json:"return_number" gorm:"type:varchar(100);not null"`
	Reason              string               `json:"reason" gorm:"type:varchar(255);not null"`
	Notes               string               `json:"notes" gorm:"type:text"`
	CreatedBy           string               `json:"created_by" gorm:"type:char(36);not null;index"`
	User                User                 `json:"-" gorm:"foreignKey:CreatedBy"`
	CreatedAt           time.Time            `gorm:"column:created_at;autoCreateTime"`
	PurchaseReturnItems []PurchaseReturnItem `gorm:"foreignKey:PurchaseReturnID"` // one purchase return has many purchase return items, relasi detail transaksi
}

func (PurchaseReturn) TableName() string {
	return "purchase_returns"
}
