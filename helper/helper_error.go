package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct error global -> definisikan field untuk segala jenis error
type AllErrors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func ErrorResponse(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, AllErrors{
		Code:    http.StatusInternalServerError,
		Message: message,
		Error:   err.Error(),
	})
}

func ErrorParsingRequestBody(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, AllErrors{
		Code:    http.StatusBadRequest,
		Message: "bad request, please check your input",
		Error:   err.Error(),
	})
}
