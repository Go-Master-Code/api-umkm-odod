package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
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

func (h *CatalogItemHandler) GetCatalogItemByID(c *gin.Context) {
	// ambil param id
	id := c.Param("id")
	ci, err := h.service.GetCatalogItemByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, ci)
}

func (h *CatalogItemHandler) CreateCatalogItem(c *gin.Context) {
	// parsing request body
	var req dto.CreateCatalogItemRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	ciDTO, err := h.service.CreateCatalogItem(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, ciDTO)
}

func (h *CatalogItemHandler) UpdateCatalogItem(c *gin.Context) {
	// parsing request body
	var req dto.UpdateCatalogItemRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")

	// update
	ciDTO, err := h.service.UpdateCatalogItem(c.Request.Context(), id, req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, ciDTO)
}

func (h *CatalogItemHandler) DeleteCatalogItem(c *gin.Context) {
	// get param id
	id := c.Param("id")

	ciDTO, err := h.service.DeleteCatalogItem(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorDeleteData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessDeleteData, ciDTO)
}
