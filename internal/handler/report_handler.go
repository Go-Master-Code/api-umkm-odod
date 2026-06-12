package handler

import (
	"fmt"
	"time"
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

func (h *ReportHandler) GetPurchaseReport(c *gin.Context) {
	var query dto.PurchaseReportQuery

	// bind query from URL
	err := c.ShouldBindQuery(&query)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	purchaseReport, err := h.service.GetPurchaseReport(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, purchaseReport)
}

func (h *ReportHandler) GetStockReport(c *gin.Context) {
	// ambil query parameter dari URL
	var query dto.StockReportQuery

	err := c.ShouldBindQuery(&query)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	helper.NormalizePagination(&query.Page, &query.Limit)

	stockReport, total, err := h.service.GetStockReport(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessGenerateStockReport(c, stockReport, int(total), query.Page, query.Limit)
}

// =========================================
// =========EXPORT REPORTS TO EXCEL=========
// =========================================
func (h *ReportHandler) ExportSalesReport(c *gin.Context) {
	var query dto.SaleReportQuery

	// bind query dari URL
	if err := c.ShouldBindQuery(&query); err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	file, err := h.service.ExportSalesReport(
		c.Request.Context(),
		query,
	)

	if err != nil {
		helper.ErrorGenerateReport(c, err)
		return
	}

	// generate file name
	fileName := fmt.Sprintf(
		"sales-report-%s.xlsx",
		time.Now().Format("20060102150405"), //yyyymmddhhmmss
	)

	c.Header(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=%s", fileName),
	)

	c.Header(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	)

	if err := file.Write(c.Writer); err != nil {
		helper.ErrorGenerateReport(c, err)
		return
	}
}

func (h *ReportHandler) ExportPurchaseReport(c *gin.Context) {
	var query dto.PurchaseReportQuery

	// bind query dari URL
	if err := c.ShouldBindQuery(&query); err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	file, err := h.service.ExportPurchaseReport(
		c.Request.Context(),
		query,
	)

	if err != nil {
		helper.ErrorGenerateReport(c, err)
		return
	}

	// generate file name
	fileName := fmt.Sprintf(
		"purchase-report-%s.xlsx",
		time.Now().Format("20060102150405"), //yyyymmddhhmmss
	)

	c.Header(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=%s", fileName),
	)

	c.Header(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	)

	if err := file.Write(c.Writer); err != nil {
		helper.ErrorGenerateReport(c, err)
		return
	}
}
