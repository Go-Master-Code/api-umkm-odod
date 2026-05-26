package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterSaleRoutes(rg *gin.RouterGroup, h *handler.SaleHandler) {
	// endpoint sale
	rg.POST("/sales", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin, constants.RoleCashier), h.CreateSale)
}
