package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterStockMovementRoutes(rg *gin.RouterGroup, h *handler.StockMovementHandler) {
	rg.POST("stock/add", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.AddStock)
	rg.POST("stock/reduce", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ReduceStock)
	rg.GET("stock/current/:itemVariantID", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin, constants.RoleCashier), h.GetCurrentStock)
	// bagian stock adjustment
	rg.POST("stock-adjustments", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateAdjustment)
}
