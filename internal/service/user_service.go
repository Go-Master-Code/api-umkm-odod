package service

import (
	"context"
	"errors"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/repository"
	"umkm-odod/internal/utils/crypto"

	"github.com/google/uuid"
)

// interface
type UserService interface {
	GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) // untuk super admin melihat semua user dari setiap tenant
	GetUsersByTenant(ctx context.Context, username string) ([]dto.UserResponse, error)
	GetUserByID(ctx context.Context, id string) (dto.UserResponse, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error)
	UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error)
	DeleteUser(ctx context.Context, id string) (dto.UserResponse, error)
	Login(ctx context.Context, username, password string) (dto.UserResponse, error)
}

// struct implementasi
type userService struct {
	repo repository.UserRepository
}

// constructor
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// struct method
func (s *userService) GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersDTO := helper.ConvertToDTOUserPlural(users)
	return usersDTO, nil
}

func (s *userService) GetUsersByTenant(ctx context.Context, username string) ([]dto.UserResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	users, err := s.repo.GetUsersByTenant(ctx, tenantID, username)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	usersDTO := helper.ConvertToDTOUserPlural(users)
	return usersDTO, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (dto.UserResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	user, err := s.repo.GetUserByID(ctx, tenantID, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOUserSingle(user)
	return userDTO, nil
}

func (s *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// convert request to model
	user := model.User{
		ID:       uuid.NewString(),
		TenantID: tenantID,
		RoleID:   req.RoleID,
		FullName: req.FullName,
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		IsActive: true,
	}

	err := s.repo.CreateUser(ctx, &user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// get user by id untuk melakukan preload relasi
	newUser, err := s.repo.GetUserByID(ctx, tenantID, user.ID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOUserSingle(newUser)
	return userDTO, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	// inisialisasi map
	var updateMap = map[string]any{}

	if req.FullName != nil {
		updateMap["full_name"] = req.FullName
	}
	if req.IsActive != nil {
		updateMap["is_active"] = req.IsActive
	}
	if req.Phone != nil {
		updateMap["phone"] = req.Phone
	}
	if req.RoleID != nil {
		updateMap["role_id"] = req.RoleID
	}
	if req.Username != nil {
		updateMap["username"] = req.Username
	}

	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// update ke repo
	err := s.repo.UpdateUser(ctx, tenantID, id, updateMap)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// get data by id untuk preload data relasi
	user, err := s.repo.GetUserByID(ctx, tenantID, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOUserSingle(user)
	return userDTO, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) (dto.UserResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data by id untuk ditampilkan di response
	user, err := s.repo.GetUserByID(ctx, tenantID, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	err = s.repo.DeleteUser(ctx, tenantID, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOUserSingle(user)
	return userDTO, nil
}

func (s *userService) Login(ctx context.Context, username, password string) (dto.UserResponse, error) {
	// get data user dulu
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return dto.UserResponse{}, errors.New("username not found")
	}

	// bandingkan password yang diinput dengan password di db
	valid := crypto.CheckPassword(password, user.Password)
	if !valid {
		return dto.UserResponse{}, errors.New("wrong password")
	}

	// jika password valid, return data
	userDTO := helper.ConvertToDTOUserSingle(user)
	return userDTO, nil

}
