package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// langsung struct implementasi
type TenantHandler struct {
	service service.TenantService
}

// constructor
func NewTenantHandler(service service.TenantService) *TenantHandler {
	return &TenantHandler{
		service: service,
	}
}

// struct method

func (h *TenantHandler) GetTenants(c *gin.Context) {
	// coba dulu ambil query param name (cek ada atau tidak)
	name := c.Query("name")

	tenants, err := h.service.GetTenants(c.Request.Context(), name)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, tenants)
}

func (h *TenantHandler) GetTenantByID(c *gin.Context) {
	// ambil param
	id := c.Param("id")

	tenant, err := h.service.GetTenantByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, tenant)
}

func (h *TenantHandler) CreateTenant(c *gin.Context) {
	// parsing json body
	var req dto.CreateTenantRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	tenantDTO, err := h.service.CreateTenant(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, tenantDTO)
}

func (h *TenantHandler) UpdateTenant(c *gin.Context) {
	// parsing request body
	var req dto.UpdateTenantRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")

	tenantDTO, err := h.service.UpdateTenant(c.Request.Context(), id, req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, tenantDTO)
}

func (h *TenantHandler) DeleteTenant(c *gin.Context) {
	// ambil id dari param
	id := c.Param("id")

	response, err := h.service.DeleteTenant(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorDeleteData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessDeleteData, response)
}
