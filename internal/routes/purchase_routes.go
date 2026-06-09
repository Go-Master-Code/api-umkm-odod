package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPurchaseRoutes(rg *gin.RouterGroup, h handler.PurchaseHandler) {
	// endpoint purchase
	rg.GET("/purchase", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetAllPurchases)
	rg.POST("/purchase", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreatePurchase)
	rg.GET("/purchase/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetPurchaseByID)
}
