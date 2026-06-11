package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(rg *gin.RouterGroup, h *handler.DashboardHandler) {
	rg.GET("dashboard/summary", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSummary)
}
