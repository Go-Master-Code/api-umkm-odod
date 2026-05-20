package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(rg *gin.RouterGroup, h *handler.RoleHandler) {
	// endpoint roles
	rg.GET("/roles", h.GetRoles)
	rg.GET("/roles/:id", h.GetRoleByID)
	rg.POST("/roles/", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateRole)
	rg.PUT("/roles/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateRole)
	rg.DELETE("//roles/:id", h.DeleteRole)
}
