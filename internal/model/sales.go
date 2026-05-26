package model

import "time"

type Sale struct {
	ID             string     `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID       string     `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant         Tenant     `json:"-" gorm:"foreignKey:TenantID"`
	InvoiceNumber  string     `json:"invoice_number" gorm:"type:varchar(100);not null;uniqueIndex"` // harus unique
	CustomerName   string     `json:"customer_name" gorm:"type:varchar(150)"`
	CashierID      string     `json:"cashier_id" gorm:"type:char(36);not null;index"`
	Cashier        User       `json:"-" gorm:"foreignKey:CashierID"`
	Subtotal       float64    `json:"subtotal" gorm:"type:decimal(18,2);not null"`
	DiscountAmount float64    `json:"discount_amount" gorm:"type:decimal(18,2);not null"`
	TaxAmount      float64    `json:"tax_amount" gorm:"type:decimal(18,2);not null"`
	GrandTotal     float64    `json:"grand_total" gorm:"type:decimal(18,2);not null"`
	PaymentMethod  string     `json:"payment_method" gorm:"type:varchar(10);not null,index"`
	PaymentStatus  string     `json:"payment_status" gorm:"type:varchar(50);not null;index"` // berguna untuk filter mana yang PAID UNPAID PARTIAL VOID REFUNDED
	Notes          string     `json:"notes" gorm:"type:text"`
	SaleItems      []SaleItem `json:"-" gorm:"foreignKey:SaleID"` // one sale has many sale items, relasi detail transaksi
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;autoUpdateTime"`
}

func (Sale) TableName() string {
	return "sales"
}
