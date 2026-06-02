package model

type Purchase struct {
	ID             string   `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID       string   `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant         Tenant   `json:"-" gorm:"foreignKey:TenantID"`
	PurchaseNumber string   `json:"purchase_number" gorm:"type:varchar(100);not null"`
	SupplierID     string   `json:"supplier_id" gorm:"type:char(36);not null;index"`
	Supplier       Supplier `json:"-" gorm:"foreignKey:SupplierID"`
	InvoiceNumber  string   `json:"invoice_number" gorm:"type:varchar(100);not null"`
	Subtotal       float64  `json:"subtotal" gorm:"type:decimal(18,2);not null"`
	DiscountAmount float64  `json:"discount_amount" gorm:"type:decimal(18,2);not null"`
	TaxAmount      float64  `json:"tax_amount" gorm:"type:decimal(18,2);not null"`
	GrandTotal     float64  `json:"grand_total" gorm:"type:decimal(18,2);not null"`
	Notes          string   `json:"notes" gorm:"type:text"`
	CreatedBy      string   `json:"created_by" gorm:"type:char(36);not null;index"`
	User           User     `json:"-" gorm:"foreignKey:CreatedBy"`
}

func (Purchase) TableName() string {
	return "purchases"
}
