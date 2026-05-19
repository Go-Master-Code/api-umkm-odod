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

type LoginSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Token   string `json:"token"`
}

func SuccessResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, AllSuccess{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func SuccessLogin(c *gin.Context, data any, token string) {
	c.JSON(http.StatusOK, LoginSuccess{
		Code:    http.StatusOK,
		Message: "login success",
		Data:    data,
		Token:   token,
	})
}
