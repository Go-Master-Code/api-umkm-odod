package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct impelementasi
type DashboardHandler struct {
	service service.DashboardService
}

// constructor
func NewDashboardHandler(service service.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

// struct method
func (h *DashboardHandler) GetSummary(c *gin.Context) {
	dashboardSummaryDTO, err := h.service.GetSummary(c.Request.Context())
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, dashboardSummaryDTO)
}

func (h *DashboardHandler) GetDailySalesChart(c *gin.Context) {
	dailySalesChart, err := h.service.GetDailySalesChart(c.Request.Context())
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, dailySalesChart)
}

func (h *DashboardHandler) GetDailyPurchaseChart(c *gin.Context) {
	dailyPurchaseChart, err := h.service.GetDailyPurchaseChart(c.Request.Context())
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, dailyPurchaseChart)
}

func (h *DashboardHandler) GetTopSellingProducts(c *gin.Context) {
	topSellingProducts, err := h.service.GetTopSellingProducts(c.Request.Context())
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, topSellingProducts)
}
