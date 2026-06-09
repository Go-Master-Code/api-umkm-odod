package dto

import "time"

type PurchaseResponse struct {
	ID             string                 `json:"id"`
	TenantID       string                 `json:"tenant_id"`
	TenantName     string                 `json:"tenant_name"`
	PurchaseNumber string                 `json:"purchase_number"`
	SupplierID     string                 `json:"supplier_id"`
	SupplierName   string                 `json:"supplier_name"`
	InvoiceNumber  string                 `json:"invoice_number"`
	Subtotal       float64                `json:"subtotal"`
	DiscountAmount float64                `json:"discount_amount"`
	TaxAmount      float64                `json:"tax_amount"`
	GrandTotal     float64                `json:"grand_total"`
	Notes          string                 `json:"notes"`
	CreatedBy      string                 `json:"created_by"`
	CreatedByName  string                 `json:"created_by_name"`
	CreatedAt      time.Time              `json:"created_at"`
	Items          []PurchaseItemResponse `json:"items"` // tampilkan sale item sebagai nested slice
}

// create
type CreatePurchaseRequest struct {
	// PurchaseNumber string `json:"invoice_number" binding:"required,max=100"` harusnya generate otomatis backend karena unique, anti manipulasi, sequence aman
	// CreatedBy      string  `json:"cashier_id" binding:"required,uuid"` dari session frontend
	// TaxAmount      float64 `json:"tax_amount" binding:"required"` idealnya dihitung oleh backend juga
	// Subtotal       float64 `json:"subtotal" binding:"required"` best practice: tidak dikirim oleh frontend, harus dihitung backend
	// GrandTotal     float64 `json:"grand_total" binding:"required"` best practice: tidak dikirim oleh frontend, harus dihitung backend
	SupplierID     string                            `json:"supplier_id" binding:"required,uuid"`
	InvoiceNumber  string                            `json:"invoice_number" binding:"omitempty,max=100"`
	DiscountAmount float64                           `json:"discount_amount" binding:"gte=0"`
	Notes          string                            `json:"notes" binding:"omitempty,max=500"`
	Items          []CreatePurchaseItemDetailRequest `json:"items" binding:"required,min=1,dive"` // tabel detil sales berisi item yang dijual
}

// update sebaiknya tidak ada karena transaksi purchase harusnya bersifat immutable, tidak bisa diedit karena akan merusak riwayat stok pula

// detil sale item pada transaksi
type CreatePurchaseItemDetailRequest struct {
	ItemVariantID  string  `json:"item_variant_id" binding:"required,uuid"`
	CostPrice      float64 `json:"cost_price" binding:"required,gte=0"`
	Qty            float64 `json:"qty" binding:"required,gt=0"`
	DiscountAmount float64 `json:"discount_amount" binding:"gte=0"`
}

// response sale item
type PurchaseItemResponse struct {
	ID                  string    `json:"id"`
	TenantID            string    `json:"tenant_id"`
	TenantName          string    `json:"tenant_name"`
	PurchaseID          string    `json:"purchase_id"`
	ItemVariantID       string    `json:"item_variant_id"`
	ItemVariantName     string    `json:"item_variant_name"`
	ItemNameSnapshot    string    `json:"item_name_snapshot"`
	VariantNameSnapshot string    `json:"variant_name_snapshot"`
	SKUSnapshot         string    `json:"sku_snapshot"`
	Qty                 float64   `json:"qty"`
	CostPrice           float64   `json:"cost_price"`
	DiscountAmount      float64   `json:"discount_amount"`
	Subtotal            float64   `json:"subtotal"`
	CreatedAt           time.Time `json:"created_at"`
}

// query params get all purchases
type GetAllPurchasesQuery struct {
	Page   int    `form:"page"`   // nomor halaman
	Limit  int    `form:"limit"`  // jumlah data per halaman
	Search string `form:"search"` // untuk search INV atau supplier
}
