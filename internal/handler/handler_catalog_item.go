package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type CatalogItemHandler struct {
	service service.CatalogItemService
}

// constructor
func NewCatalogItemHandler(service service.CatalogItemService) *CatalogItemHandler {
	return &CatalogItemHandler{
		service: service,
	}
}

// struct method
func (h *CatalogItemHandler) GetCatalogItems(c *gin.Context) {
	// coba ambil query param name
	name := c.Query("name")

	ciDTO, err := h.service.GetCatalogItems(c.Request.Context(), name)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, ciDTO)
}
