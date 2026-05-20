package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCatalogItemRoutes(rg *gin.RouterGroup, h *handler.CatalogItemHandler) {
	// endpoint catalog items
	rg.GET("/catalog_items", h.GetCatalogItems)
	rg.GET("/catalog_items/:id", h.GetCatalogItemByID)
	rg.POST("/catalog_items", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateCatalogItem)
	rg.PUT("/catalog_items/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateCatalogItem)
	rg.DELETE("/catalog_items/:id", h.DeleteCatalogItem)
}
