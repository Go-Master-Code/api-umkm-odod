package routes

import (
	"umkm-odod/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterActivityLogRoutes(rg *gin.RouterGroup, h *handler.ActivityLogHandler) {
	rg.GET("/log", h.GetAllActivityLog) //sementara dibuat tanpa middleware dulu
}
