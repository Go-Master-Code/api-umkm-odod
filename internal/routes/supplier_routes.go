package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterSupplierRoutes(rg *gin.RouterGroup, h *handler.SupplierHandler) {
	// endpoint supplier
	rg.GET("/suppliers", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSuppliers)
	rg.GET("/suppliers/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSupplierByID)
	rg.POST("/suppliers", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateSupplier)
	rg.PUT("/suppliers/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateSupplier)
	rg.DELETE("/suppliers/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.DeleteSupplier)
	// rg.GET("/users/me", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetUsersByTenant) // lihat ada user siapa saja di tenant saya
	// rg.GET("/users/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetUserByID)
	// rg.POST("/users", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateUser)
	// rg.PUT("/users/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateUser)
	// rg.DELETE("/users/:id", middleware.AuthRole(constants.RoleOwner), h.DeleteUser)
}
