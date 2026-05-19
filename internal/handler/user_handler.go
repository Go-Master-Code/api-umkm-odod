package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"
	"umkm-odod/internal/utils/crypto"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct
type UserHandler struct {
	service service.UserService
}

// constructor
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// struct method
func (h *UserHandler) GetUsers(c *gin.Context) {
	// ambil query username jika ada
	username := c.Query("username")

	users, err := h.service.GetUsers(c.Request.Context(), username)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	// ambil param id
	id := c.Param("id")

	user, err := h.service.GetUserByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// parsing request body
	var req dto.CreateUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// hash password yang diinput user
	hashPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorHashPassword, err)
		return
	}

	// ganti password pada request pakai yang sudah di hash
	req.Password = hashPassword

	newUser, err := h.service.CreateUser(c.Request.Context(), req)

	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, newUser)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	// parsing request body
	var req dto.UpdateUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")

	userDTO, err := h.service.UpdateUser(c.Request.Context(), id, req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, userDTO)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// ambil param id
	id := c.Param("id")

	userDTO, err := h.service.DeleteUser(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorDeleteData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessDeleteData, userDTO)
}
