package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPurchaseReturnRoutes(rg *gin.RouterGroup, h handler.PurchaseReturnHandler) {
	// endpoint purchase return
	rg.GET("/purchase_return", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetAllPurchaseReturns)
	rg.GET("/purchase_return/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetPurchaseReturnByID)
	rg.POST("/purchase_return", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreatePurchaseReturn)
}
