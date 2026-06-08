package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type PurchaseHandler struct {
	service service.PurchaseService
}

// constructor
func NewPurchaseHandler(service service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{
		service: service,
	}
}

// struct method
func (h *PurchaseHandler) CreatePurchase(c *gin.Context) {
	var req dto.CreatePurchaseRequest

	// parsing request body
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// call service
	purchaseDTO, err := h.service.CreatePurchase(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, purchaseDTO)
}
