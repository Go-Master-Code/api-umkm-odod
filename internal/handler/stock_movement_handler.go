package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type StockMovementHandler struct {
	service service.StockMovementService
}

// constructor
func NewStockMovementHandler(service service.StockMovementService) *StockMovementHandler {
	return &StockMovementHandler{
		service: service,
	}
}

// struct method
func (h *StockMovementHandler) AddStock(c *gin.Context) {
	// parsing request body
	var req dto.AddStockRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	stockMovementDTO, err := h.service.AddStock(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, "stock added successfully", stockMovementDTO)
}

func (h *StockMovementHandler) ReduceStock(c *gin.Context) {
	// parsing request body
	var req dto.ReduceStockRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	stockMovementDTO, err := h.service.ReduceStock(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, "stock reduced successfully", stockMovementDTO)
}

func (h *StockMovementHandler) GetCurrentStock(c *gin.Context) {
	// ambil item variant id dari param
	itemVariantID := c.Param("itemVariantID")

	currentStockDTO, err := h.service.GetCurrentStock(c.Request.Context(), itemVariantID)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, currentStockDTO)
}

// adjustment section
func (h *StockMovementHandler) CreateAdjustment(c *gin.Context) {
	// binding request body
	var req dto.CreateStockAdjustmentRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// call service
	smDTO, err := h.service.CreateAdjustment(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	// success response
	helper.SuccessResponse(c, constants.SuccessCreateData, smDTO)
}

// stock card
func (h *StockMovementHandler) GetStockCard(c *gin.Context) {
	// ambil item variant id dari param URL
	itemVariantID := c.Param("itemVariantID")

	stockCard, err := h.service.GetStockCard(c.Request.Context(), itemVariantID)

	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, stockCard)
}
