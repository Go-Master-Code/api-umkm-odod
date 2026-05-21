package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
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

func (h *ItemVariantHandler) GetItemVariantByID(c *gin.Context) {
	// tangkap param id
	id := c.Param("id")

	ivDTO, err := h.service.GetItemVariantByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, ivDTO)
}

func (h *ItemVariantHandler) CreateItemVariant(c *gin.Context) {
	var req dto.CreateItemVariantRequest

	// parsing json body
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	ivDTO, err := h.service.CreateItemVariant(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, ivDTO)
}

func (h *ItemVariantHandler) UpdateItemVariant(c *gin.Context) {
	var req dto.UpdateItemVariantRequest

	// parsing json body
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")

	ivDTO, err := h.service.UpdateItemVariant(c.Request.Context(), id, req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, ivDTO)
}
