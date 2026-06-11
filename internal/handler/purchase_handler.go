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

func (h *PurchaseHandler) GetAllPurchases(c *gin.Context) {
	var query dto.GetAllPurchasesQuery
	err := c.ShouldBindQuery(&query) // bind semua query dari URL

	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// pagination with helper
	helper.NormalizePagination(&query.Page, &query.Limit)

	purchaseResponsesDTO, total, err := h.service.GetAllPurchases(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessGetAllPurchasesPerTenant(c, purchaseResponsesDTO, int(total), query.Page, query.Limit)
}

func (h *PurchaseHandler) GetPurchaseByID(c *gin.Context) {
	id := c.Param("id")

	purchaseDTO, err := h.service.GetPurchaseByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, purchaseDTO)
}
