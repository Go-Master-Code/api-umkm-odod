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

func (h *PurchaseReturnHandler) GetPurchaseReturnByID(c *gin.Context) {
	// get id param from URL
	id := c.Param("id")

	purchaseReturnDTO, err := h.service.GetPurchaseReturnByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, purchaseReturnDTO)
}

func (h *PurchaseReturnHandler) GetAllPurchaseReturns(c *gin.Context) {
	var query dto.GetAllPurchaseReturnsQuery
	err := c.ShouldBindQuery(&query) // bind semua query dari URL

	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// pagination with helper
	helper.NormalizePagination(&query.Page, &query.Limit)

	purchaseReturnDTO, total, err := h.service.GetAllPurchaseReturns(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessGetAllPurchaseReturnPerTenant(c, purchaseReturnDTO, int(total), query.Page, query.Limit)
}
