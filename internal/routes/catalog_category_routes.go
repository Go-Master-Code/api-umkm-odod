package routes

import (
	"umkm-odod/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCatalogCategoryRoutes(rg *gin.RouterGroup, h *handler.CatalogCategoryHandler) {
	// endpoint roles
	rg.GET("/catalog_categories", h.GetCatalogCategories)
	// rg.GET("/roles/:id", h.GetRoleByID)
	// rg.POST("/roles/", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateRole)
	// rg.PUT("/roles/:id", h.UpdateRole)
	// rg.DELETE("//roles/:id", h.DeleteRole)
}
