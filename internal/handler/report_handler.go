package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type ReportHandler struct {
	service service.ReportService
}

// constructor
func NewReportHandler(service service.ReportService) *ReportHandler {
	return &ReportHandler{
		service: service,
	}
}

// struct method
func (h *ReportHandler) GetSalesReport(c *gin.Context) {
	var query dto.SaleReportQuery

	// parsing query
	err := c.ShouldBindQuery(&query)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	salesReport, err := h.service.GetSalesReport(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, salesReport)
}
