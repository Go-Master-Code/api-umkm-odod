package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup, h *handler.ReportHandler) {
	rg.GET("/reports/sales", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSalesReport)
	rg.GET("/reports/purchase", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetPurchaseReport)
}
