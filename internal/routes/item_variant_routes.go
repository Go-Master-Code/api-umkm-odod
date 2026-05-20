package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterItemVariantRoutes(rg *gin.RouterGroup, h *handler.ItemVariantHandler) {
	rg.GET("/item_variants", h.GetItemVariants)
	rg.GET("/item_variants/:id", h.GetItemVariantByID)
	rg.POST("/item_variants", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateItemVariant)
}
