package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type SupplierHandler struct {
	service service.SupplierService
}

// constructor
func NewSupplierHandler(service service.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		service: service,
	}
}

// struct method
func (h *SupplierHandler) GetSuppliers(c *gin.Context) {
	// coba get query name
	name := c.Query("name")

	suppliersDTO, err := h.service.GetSuppliers(c.Request.Context(), name)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, suppliersDTO)
}

func (h *SupplierHandler) GetSupplierByID(c *gin.Context) {
	// param
	id := c.Param("id")

	supplierDTO, err := h.service.GetSupplierByID(c.Request.Context(), id)

	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, supplierDTO)
}

func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	// parsing request body
	var req dto.CreateSupplierRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	supplierDTO, err := h.service.CreateSupplier(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, supplierDTO)
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")

	// parsing request body
	var req dto.UpdateSupplierRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	supplierDTO, err := h.service.UpdateSupplier(c.Request.Context(), id, req)

	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, supplierDTO)
}

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	supplierDTO, err := h.service.DeleteSupplier(c.Request.Context(), id)

	if err != nil {
		helper.ErrorResponse(c, constants.ErrorDeleteData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessDeleteData, supplierDTO)
}
