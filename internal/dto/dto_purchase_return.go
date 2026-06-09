package dto

import "time"

type PurchaseReturnResponse struct {
	ID            string                       `json:"id"`
	TenantID      string                       `json:"tenant_id"`
	TenantName    string                       `json:"tenant_name"`
	PurchaseID    string                       `json:"purchase_id"`
	ReturnNumber  string                       `json:"return_number"`
	Reason        string                       `json:"reason"`
	Notes         string                       `json:"notes"`
	CreatedBy     string                       `json:"created_by"`
	CreatedByName string                       `json:"created_by_name"`
	CreatedAt     time.Time                    `json:"created_at"`
	Items         []PurchaseReturnItemResponse `json:"items"`
}

// create
type CreatePurchaseReturnRequest struct {
	PurchaseID string `json:"purchase_id" binding:"required,uuid"`
	// ReturnNumber string                            `json:"return_number" binding:"required,max=100"` jangan terima parsing dari frontend, harusnya digenerate backend seperti PRETURN-1749550000
	Reason string                            `json:"reason" binding:"required,max=255"`
	Notes  string                            `json:"notes" binding:"omitempty,max=500"`
	Items  []CreatePurchaseReturnItemRequest `json:"items" binding:"required,min=1,dive"`
}

// update sebaiknya tidak ada karena transaksi purchase return harusnya bersifat immutable, tidak bisa diedit karena akan merusak riwayat stok

type PurchaseReturnItemResponse struct {
	ID               string    `json:"id"`
	TenantID         string    `json:"tenant_id"`
	PurchaseReturnID string    `json:"purchase_return_id"`
	ItemVariantID    string    `json:"item_variant_id"`
	ItemVariantName  string    `json:"item_variant_name"`
	Qty              float64   `json:"qty"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
}

// create detil (items) for purchase return
type CreatePurchaseReturnItemRequest struct {
	// PurchaseReturnID string  `json:"purchase_return_id" binding:"required,uuid"` PurchaseReturnID baru diketahui setelah header berhasil dibuat.
	ItemVariantID string  `json:"item_variant_id" binding:"required,uuid"`
	Qty           float64 `json:"qty" binding:"required,gt=0"`
	Notes         string  `json:"notes" binding:"omitempty,max=500"`
}

// query params get all purchases
type GetAllPurchaseReturnsQuery struct {
	Page   int    `form:"page"`   // nomor halaman
	Limit  int    `form:"limit"`  // jumlah data per halaman
	Search string `form:"search"` // untuk search search return_number, purchase_number atau reason
}
