package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type ItemVariantHandler struct {
	service service.ItemVariantService
}

// constructor
func NewItemVariantHandler(service service.ItemVariantService) *ItemVariantHandler {
	return &ItemVariantHandler{
		service: service,
	}
}

// struct method
func (h *ItemVariantHandler) GetItemVariants(c *gin.Context) {
	// tangkap query name
	name := c.Query("name")

	ivDTO, err := h.service.GetItemVariants(c.Request.Context(), name)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, ivDTO)
}
