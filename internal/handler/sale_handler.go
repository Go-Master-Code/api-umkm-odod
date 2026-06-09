package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// langsung struct impelementasi, no interface
type SaleHandler struct {
	service service.SaleService
}

// constructor
func NewSaleHandler(service service.SaleService) *SaleHandler {
	return &SaleHandler{
		service: service,
	}
}

// struct method
func (h *SaleHandler) CreateSale(c *gin.Context) {
	// parsing request body
	var req dto.CreateSaleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// call service
	saleDTO, err := h.service.CreateSale(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	// success response
	helper.SuccessResponse(c, constants.SuccessCreateData, saleDTO)
}

func (h *SaleHandler) GetAllSales(c *gin.Context) {
	// parsing request body
	var query dto.GetAllSalesQuery
	err := c.ShouldBindQuery(&query) // bind semua query dari URL
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// pagination with helper
	helper.NormalizePagination(&query.Page, &query.Limit)

	salesResponseDTO, total, err := h.service.GetAllSales(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	// jika sukses
	helper.SuccessGetAllSalesPerTenant(c, salesResponseDTO, int(total), query.Page, query.Limit)
}

func (h *SaleHandler) GetSaleByID(c *gin.Context) {
	// ambil param id
	id := c.Param("id")

	saleDTO, err := h.service.GetSaleByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, saleDTO)
}
