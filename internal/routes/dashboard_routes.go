package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(rg *gin.RouterGroup, h *handler.DashboardHandler) {
	rg.GET("/dashboard/summary", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSummary)
	rg.GET("/dashboard/chart/sales", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetDailySalesChart)
	rg.GET("/dashboard/chart/purchase", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetDailyPurchaseChart)
	rg.GET("/dashboard/chart/top-selling-products", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetTopSellingProducts)
}
