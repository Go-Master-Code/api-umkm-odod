package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
	// endpoint user
	rg.GET("/users", middleware.AuthRole(constants.RoleSuperAdmin), h.GetAllUsers)                             // hanya super admin yang bisa liat semua users
	rg.GET("/tenant-users", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetUsersByTenant) // lihat ada user siapa saja di tenant saya
	rg.GET("/users/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetUserByID)
	rg.POST("/users", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateUser)
	rg.PUT("/users/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateUser)
	rg.DELETE("/users/:id", middleware.AuthRole(constants.RoleOwner), h.DeleteUser)
	// lihat profile dari user yang aktif
	rg.GET("/users/me", h.GetProfile)
	rg.PUT("/users/me", h.UpdateProfile)
	rg.POST("/users/change-password", h.ChangePassword)
}
