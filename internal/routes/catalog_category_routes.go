package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCatalogCategoryRoutes(rg *gin.RouterGroup, h *handler.CatalogCategoryHandler) {
	// endpoint roles
	rg.GET("/catalog_categories", h.GetCatalogCategories)
	rg.GET("/catalog_categories/:id", h.GetCatalogCategoryByID)
	rg.POST("/catalog_categories/", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateCatalogCategory)
	rg.PUT("/catalog_categories/:id", h.UpdateCatalogCategory)
	// rg.DELETE("//roles/:id", h.DeleteRole)
}
