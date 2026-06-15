package model

import "time"

type ActivityLog struct {
	ID              string    `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID        string    `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant          Tenant    `json:"-" gorm:"foreignKey:TenantID"`
	UserID          string    `json:"user_id" gorm:"type:char(36);not null;index"`
	User            User      `json:"-" gorm:"foreignKey:UserID"`
	Module          string    `json:"module" gorm:"type:varchar(50);not null;index"`
	Action          string    `json:"action" gorm:"type:varchar(50);not null;index"`
	Description     string    `json:"description" gorm:"type:varchar(255);not null"`
	ReferenceID     string    `json:"reference_id" gorm:"type:char(36)"`         // Referensi teknis ke record database
	ReferenceNumber string    `json:"reference_number" gorm:"type:varchar(100)"` // Referensi yang mudah dibaca manusia
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}
