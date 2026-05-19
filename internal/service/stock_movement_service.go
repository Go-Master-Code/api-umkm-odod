package service

import (
	"umkm-odod/internal/dto"
	"umkm-odod/internal/payloads"
)

// interface
type StockMovementService interface {
	CreateMovement(payload payloads.CreateStockMovementPayload) dto.StockMovementResponse
}

// struct implementasi
type stockMovementService struct {
	// repo
}

// constructor
func NewStockMovementService() {

}
