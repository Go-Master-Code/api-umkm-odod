package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type PurchaseReturnHandler struct {
	service service.PurchaseReturnService
}

// constructor
func NewPurchaseReturnHandler(service service.PurchaseReturnService) *PurchaseReturnHandler {
	return &PurchaseReturnHandler{service: service}
}

// struct method
func (h *PurchaseReturnHandler) CreatePurchaseReturn(c *gin.Context) {
	// parsing request body
	var req dto.CreatePurchaseReturnRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// jika parsing request body sukses, execute service
	purchaseReturnDTO, err := h.service.CreatePurchaseReturn(c.Request.Context(), req)

	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, purchaseReturnDTO)
}
