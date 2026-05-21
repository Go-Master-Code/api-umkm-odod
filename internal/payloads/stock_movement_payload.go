package payloads

// ini adalah payload internal -> untuk komunikasi antar service
// payload internal BUKAN REQUEST dari frontend
// tidak perlu json, gorm tag, dan juga binding

// Kenapa payload internal penting?
// Karena nanti:
// sales service
// ↓
// stock service
// ↓
// create stock movement
// Frontend tidak pernah langsung membuat stock movement.

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
