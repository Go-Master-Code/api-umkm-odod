package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type UserRepository interface {
	GetUsers(ctx context.Context, username string) ([]model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error) // untuk login user
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, id string, updateMap map[string]any) error
	DeleteUser(ctx context.Context, id string) error
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
func (r *userRepository) GetUsers(ctx context.Context, username string) ([]model.User, error) {
	var users []model.User

	// query standar
	query := r.db.WithContext(ctx).Preload("Role").Preload("Tenant")

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

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
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

func (r *userRepository) UpdateUser(ctx context.Context, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.User{}).Where("id = ?", id).Updates(updateMap).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}
