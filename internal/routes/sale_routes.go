package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterSaleRoutes(rg *gin.RouterGroup, h *handler.SaleHandler) {
	// endpoint sale
	rg.GET("/sales", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetAllSales)
	rg.POST("/sales", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin, constants.RoleCashier), h.CreateSale)
	rg.GET("/sales/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSaleByID)
}
