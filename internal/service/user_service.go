package service

import (
	"context"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/repository"

	"github.com/google/uuid"
)

// interface
type UserService interface {
	GetUsers(ctx context.Context, username string) ([]dto.UserResponse, error)
	GetUserByID(ctx context.Context, id string) (dto.UserResponse, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error)
	UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error)
	DeleteUser(ctx context.Context, id string) (dto.UserResponse, error)
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
func (s *userService) GetUsers(ctx context.Context, username string) ([]dto.UserResponse, error) {
	users, err := s.repo.GetUsers(ctx, username)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	usersDTO := helper.ConvertToDTOUserPlural(users)
	return usersDTO, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (dto.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOUserSingle(user)
	return userDTO, nil
}

func (s *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
	// payload sementara
	tenantID := "f27e441f-5385-4b8d-b2e2-88b8615a4634"

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
	newUser, err := s.repo.GetUserByID(ctx, user.ID)
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

	// update ke repo
	err := s.repo.UpdateUser(ctx, id, updateMap)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// get data by id untuk preload data relasi
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOUserSingle(user)
	return userDTO, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) (dto.UserResponse, error) {
	// get data by id untuk ditampilkan di response
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	err = s.repo.DeleteUser(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOUserSingle(user)
	return userDTO, nil
}
