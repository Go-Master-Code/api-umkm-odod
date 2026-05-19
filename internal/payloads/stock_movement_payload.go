package payloads

// tidak perlu json dan gorm tag
type CreateStockMovementPayload struct {
	TenantID      string
	ItemVariantID string
	MovementType  string
	Qty           float64
	ReferenceType string
	ReferenceID   string
	Notes         string
	CreatedBy     string
}
