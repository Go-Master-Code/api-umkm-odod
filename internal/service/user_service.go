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
	"golang.org/x/crypto/bcrypt"
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
	// service endpoint me
	GetProfile(ctx context.Context) (dto.ProfileResponse, error)
	UpdateProfile(ctx context.Context, req dto.UpdateProfileRequest) (dto.ProfileResponse, error)
	ChangePassword(ctx context.Context, req dto.ChangePasswordRequest) (dto.ProfileResponse, error)
	// update last login at
	UpdateLastLoginAt(ctx context.Context, userID string) error
}

// struct implementasi
type userService struct {
	repo repository.UserRepository
	// log
	activityLogService ActivityLogService // jangan pakai package service, karena kedua file ini ada di dalam package yang sama (service)
}

// constructor
func NewUserService(repo repository.UserRepository, activityLogService ActivityLogService) UserService {
	return &userService{
		repo:               repo,
		activityLogService: activityLogService,
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

func (s *userService) GetProfile(ctx context.Context) (dto.ProfileResponse, error) {
	// user ID ambil dari jwt
	userID := ctx.Value(constants.ContextUserID).(string)

	user, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// convert model to dto
	userDTO := helper.ConvertToDTOProfileUser(user)
	return userDTO, nil
}

func (s *userService) UpdateProfile(ctx context.Context, req dto.UpdateProfileRequest) (dto.ProfileResponse, error) {
	// get userID from jwt
	userID := ctx.Value(constants.ContextUserID).(string)

	// get data model user
	user, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// isi field fullname dan phone dengan input request
	user.FullName = req.FullName
	user.Phone = req.Phone

	err = s.repo.UpdateProfile(ctx, user)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// tambahkan activity log
	err = s.activityLogService.CreateActivityLog(ctx, "USER", "UPDATE", "Update Own Profile", userID, user.Username)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// get data lagi agar terupdate
	profile, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// convert model to dto
	profileDTO := helper.ConvertToDTOProfileUser(profile)

	return profileDTO, nil
}

func (s *userService) ChangePassword(ctx context.Context, req dto.ChangePasswordRequest) (dto.ProfileResponse, error) {
	// get userID from jwt
	userID := ctx.Value(constants.ContextUserID).(string)
	user, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// cek password lama
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
	if err != nil {
		return dto.ProfileResponse{}, errors.New("old password is incorrect")
	}

	// bandingkan newPassword dengan confirmPassword (harus sama)
	if req.NewPassword != req.ConfirmPassword {
		return dto.ProfileResponse{}, errors.New("confirm password doesn't match")
	}

	// hash password baru setelah newPassword = confirmPassword
	newPassword, err := crypto.HashPassword(req.NewPassword)
	if err != nil {
		return dto.ProfileResponse{}, errors.New("failed to hash new password")
	}

	// jika newPassword dan confirmPassword sudah sama, ubah password (hasil hash) ke DB untuk user tersebut
	user.Password = newPassword

	err = s.repo.ChangePassword(ctx, user)
	if err != nil {
		return dto.ProfileResponse{}, errors.New("failed to update password")
	}

	// setelah berhasil change password, buat log nya
	err = s.activityLogService.CreateActivityLog(ctx, "USER", "CHANGE PASSWORD", "Change own password", user.ID, user.Username)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// convert model to dto untuk ditampilkan
	userDTO := helper.ConvertToDTOProfileUser(user)
	return userDTO, nil
}

func (s *userService) UpdateLastLoginAt(ctx context.Context, userID string) error {
	err := s.repo.UpdateLastLoginAt(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
