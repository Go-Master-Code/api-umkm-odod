package routes

import (
	"umkm-odod/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterItemVariantRoutes(rg *gin.RouterGroup, h *handler.ItemVariantHandler) {
	rg.GET("/item_variant", h.GetItemVariants)
}
