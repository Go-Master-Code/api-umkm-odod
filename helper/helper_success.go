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

type GetAllSalesOrPurchasePerTenantSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Total   int    `json:"total"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
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

func SuccessGetAllSalesPerTenant(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchasePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success get all sales data",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

func SuccessGetAllPurchasesPerTenant(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchasePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success get all purchases data",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}
