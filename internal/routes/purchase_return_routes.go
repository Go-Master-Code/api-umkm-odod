package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPurchaseReturnRoutes(rg *gin.RouterGroup, h handler.PurchaseReturnHandler) {
	// endpoint purchase return
	rg.POST("/purchase_return", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreatePurchaseReturn)
}
