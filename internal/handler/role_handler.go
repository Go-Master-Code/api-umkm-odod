package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type RoleHandler struct {
	service service.RoleService
}

// constructor
func NewRoleHandler(service service.RoleService) *RoleHandler {
	return &RoleHandler{
		service: service,
	}
}

// struct method
func (h *RoleHandler) GetRoles(c *gin.Context) {
	// coba ambil query param name
	name := c.Query("name")

	roles, err := h.service.GetRoles(c.Request.Context(), name)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, roles)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	// ambil param id
	id := c.Param("id")

	role, err := h.service.GetRoleByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, role)
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	// parsing json request body
	var req dto.CreateRoleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	roleDTO, err := h.service.CreateRole(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, roleDTO)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	// parsing json request body
	var req dto.UpdateRoleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil id role dari param
	id := c.Param("id")

	// update service
	roleDTO, err := h.service.UpdateRole(c.Request.Context(), id, req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, roleDTO)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	// get param id
	id := c.Param("id")

	roleDTO, err := h.service.DeleteRole(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorDeleteData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessDeleteData, roleDTO)
}
