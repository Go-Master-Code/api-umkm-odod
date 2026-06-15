package repository

import (
	"context"
	"time"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)                                        // method untuk super admin bisa lihat semua user dari semua tenant
	GetUsersByTenant(ctx context.Context, tenantID string, username string) ([]model.User, error) // admin n owner hanya boleh lihat user dari tenant sendiri, jadi harus tenant isolation
	GetUserByID(ctx context.Context, tenantID string, id string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error) // untuk login user
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	DeleteUser(ctx context.Context, tenantID string, id string) error
	// untuk endpoint me
	GetProfile(ctx context.Context, userID string) (*model.User, error)
	UpdateProfile(ctx context.Context, user *model.User) error
	ChangePassword(ctx context.Context, user *model.User) error
	// update last login at
	UpdateLastLoginAt(ctx context.Context, userID string) error
}

// struct implementasi
type userRepository struct {
	db *gorm.DB
}

// constructor
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// struct method
func (r *userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Preload("Role").Preload("Tenant").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUsersByTenant(ctx context.Context, tenantID string, username string) ([]model.User, error) {
	var users []model.User

	// query standar
	query := r.db.WithContext(ctx).Preload("Role").Preload("Tenant").Where("tenant_id = ?", tenantID)

	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}

	// search ke db
	err := query.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, tenantID string, id string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Preload("Tenant").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Preload("Tenant").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) UpdateUser(ctx context.Context, tenantID string, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.User{}).Where("id = ?", id).Updates(updateMap).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, tenantID string, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *userRepository) GetProfile(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Preload("Role").Preload("Tenant").First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateProfile(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) ChangePassword(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) UpdateLastLoginAt(ctx context.Context, userID string) error {
	// var now untuk update last login at
	now := time.Now()

	return r.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", userID).Update("last_login_at", now).Error
}
