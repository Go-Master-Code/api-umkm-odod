package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// definisikan struct global untuk success message
type AllSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, AllSuccess{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}
