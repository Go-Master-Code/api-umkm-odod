package routes

import (
	"umkm-odod/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCatalogItemRoutes(rg *gin.RouterGroup, h *handler.CatalogItemHandler) {
	rg.GET("/catalog_item", h.GetCatalogItems)
}
