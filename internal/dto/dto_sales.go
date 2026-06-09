package dto

import "time"

type SaleResponse struct {
	ID             string             `json:"id"`
	TenantID       string             `json:"tenant_id"`
	TenantName     string             `json:"tenant_name"`
	InvoiceNumber  string             `json:"invoice_number"`
	CustomerName   string             `json:"customer_name"`
	CashierID      string             `json:"cashier_id"`
	CashierName    string             `json:"cashier_name"`
	Subtotal       float64            `json:"subtotal"`
	DiscountAmount float64            `json:"discount_amount"`
	TaxAmount      float64            `json:"tax_amount"`
	GrandTotal     float64            `json:"grand_total"`
	PaymentMethod  string             `json:"payment_method"`
	PaymentStatus  string             `json:"payment_status"` // berguna untuk filter mana yang PAID UNPAID PARTIAL VOID REFUNDED
	Notes          string             `json:"notes"`
	CreatedAt      time.Time          `json:"created_at"`
	Items          []SaleItemResponse `json:"items"` // tampilkan sale item sebagai nested slice
}

// create
type CreateSaleRequest struct {
	// InvoiceNumber string `json:"invoice_number" binding:"required,max=100"` harusnya generate otomatis backend karena unique, anti manipulasi, sequence aman
	// CashierID      string  `json:"cashier_id" binding:"required,uuid"` dari session frontend
	// TaxAmount      float64 `json:"tax_amount" binding:"required"` idealnya dihitung oleh backend juga
	// Subtotal       float64 `json:"subtotal" binding:"required"` best practice: tidak dikirim oleh frontend, harus dihitung backend
	// GrandTotal     float64 `json:"grand_total" binding:"required"` best practice: tidak dikirim oleh frontend, harus dihitung backend
	CustomerName   string                        `json:"customer_name" binding:"omitempty,max=150"`
	DiscountAmount float64                       `json:"discount_amount" binding:"gte=0"`
	PaymentMethod  string                        `json:"payment_method" binding:"required,oneof=CASH QRIS TRANSFER DEBIT CREDIT"`   // tipe enum
	PaymentStatus  string                        `json:"payment_status" binding:"required,oneof=PAID UNPAID PARTIAL VOID REFUNDED"` // berguna untuk filter mana yang PAID UNPAID PARTIAL VOID REFUNDED
	Notes          string                        `json:"notes" binding:"omitempty,max=500"`
	Items          []CreateSaleItemDetailRequest `json:"items" binding:"required,min=1,dive"` // tabel detil sales berisi item yang dijual
}

/// update sebaiknya tidak ada karena transaksi sales harusnya bersifat immutable, tidak bisa diedit karena akan merusak riwayat stok pula

// detil sale item pada transaksi
type CreateSaleItemDetailRequest struct {
	ItemVariantID  string  `json:"item_variant_id" binding:"required,uuid"`
	Qty            float64 `json:"qty" binding:"required,gt=0"`
	DiscountAmount float64 `json:"discount_amount" binding:"gte=0"`
}

// response sale item
type SaleItemResponse struct {
	ID         string `json:"id"`
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	SaleID     string `json:"sale_id"`
	// InvoiceNumber       string    `json:"invoice_number"` ga usah pake karena sudah muncul di master sale
	ItemVariantID       string    `json:"item_variant_id"`
	ItemVariantName     string    `json:"item_variant_name"`
	ItemNameSnapshot    string    `json:"item_name_snapshot"`
	VariantNameSnapshot string    `json:"variant_name_snapshot"`
	SKUSnapshot         string    `json:"sku_snapshot"`
	Qty                 float64   `json:"qty"`
	UnitPrice           float64   `json:"unit_price"`
	DiscountAmount      float64   `json:"discount_amount"`
	Subtotal            float64   `json:"subtotal"`
	CreatedAt           time.Time `json:"created_at"`
}

// query params get all sales
type GetAllSalesQuery struct {
	Page          int    `form:"page"`           // nomor halaman
	Limit         int    `form:"limit"`          // jumlah data per halaman
	Search        string `form:"search"`         // untuk search INV atau customer
	PaymentStatus string `form:"payment_status"` // filter payment status
}
