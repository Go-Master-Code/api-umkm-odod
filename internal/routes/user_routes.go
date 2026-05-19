package routes

import (
	"umkm-odod/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
	// endpoint user
	rg.GET("/users", h.GetUsers)
	rg.GET("/users/:id", h.GetUserByID)
	rg.POST("/users", h.CreateUser)
	rg.PUT("/users/:id", h.UpdateUser)
	rg.DELETE("/users/:id", h.DeleteUser)
}
