package handler

import (
	"umkm-odod/auth"
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
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers(c.Request.Context())
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, users)
}

func (h *UserHandler) GetUsersByTenant(c *gin.Context) {
	// ambil query username jika ada
	username := c.Query("username")

	users, err := h.service.GetUsersByTenant(c.Request.Context(), username)
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

func (h *UserHandler) Login(c *gin.Context) {
	// parsing request body
	var req dto.LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	user, err := h.service.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil { // gagal login
		helper.ErrorResponse(c, constants.ErrorLoginInvalid, err)
		return
	}

	// jika username dan password benar, generate token jwt
	token, err := auth.GenerateToken(user.ID, user.Username, user.RoleID, user.RoleName, user.TenantID)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGenerateToken, err)
		return
	}

	// setelah token dibuat, update last login
	err = h.service.UpdateLastLoginAt(c.Request.Context(), user.ID)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	// login berhasil => kirim data user sekaligus token
	helper.SuccessLogin(c, user, token)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	profile, err := h.service.GetProfile(c.Request.Context())
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, profile)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	profileDTO, err := h.service.UpdateProfile(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, profileDTO)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	profileDTO, err := h.service.ChangePassword(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, profileDTO)
}
