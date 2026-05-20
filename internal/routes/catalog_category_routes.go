package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCatalogCategoryRoutes(rg *gin.RouterGroup, h *handler.CatalogCategoryHandler) {
	// endpoint catalog categories
	rg.GET("/catalog_categories", h.GetCatalogCategories)
	rg.GET("/catalog_categories/:id", h.GetCatalogCategoryByID)
	rg.POST("/catalog_categories/", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateCatalogCategory)
	rg.PUT("/catalog_categories/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateCatalogCategory)
	rg.DELETE("/catalog_categories/:id", h.DeleteCatalogCategory)
}
