package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// langsung struct implementasi tanpa interface
type CatalogCategoryHandler struct {
	service service.CatalogCategoryService
}

// constuctor
func NewCatalogCategoryHandler(service service.CatalogCategoryService) *CatalogCategoryHandler {
	return &CatalogCategoryHandler{
		service: service,
	}
}

// struct method
func (h *CatalogCategoryHandler) GetCatalogCategories(c *gin.Context) {
	// ambil query name jika ada
	name := c.Query("name")

	ccDTO, err := h.service.GetCatalogCategories(c.Request.Context(), name)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, ccDTO)
}
