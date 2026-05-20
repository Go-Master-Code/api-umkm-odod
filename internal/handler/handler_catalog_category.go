package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
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

func (h *CatalogCategoryHandler) GetCatalogCategoryByID(c *gin.Context) {
	// ambil param id
	id := c.Param("id")
	cc, err := h.service.GetCatalogCategoryByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, cc)
}

func (h *CatalogCategoryHandler) CreateCatalogCategory(c *gin.Context) {
	// parsing request body
	var req dto.CreateCatalogCategoryRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	ccDTO, err := h.service.CreateCatalogCategory(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, ccDTO)
}

func (h *CatalogCategoryHandler) UpdateCatalogCategory(c *gin.Context) {
	// parsing request body
	var req dto.UpdateCatalogCategoryRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")

	// update
	ccDTO, err := h.service.UpdateCatalogCategory(c.Request.Context(), id, req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, ccDTO)
}
