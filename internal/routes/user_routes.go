package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
	// endpoint user
	rg.GET("/users", h.GetUsers)
	rg.GET("/users/:id", h.GetUserByID)
	rg.POST("/users", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateUser)
	rg.PUT("/users/:id", h.UpdateUser)
	rg.DELETE("/users/:id", h.DeleteUser)
}
