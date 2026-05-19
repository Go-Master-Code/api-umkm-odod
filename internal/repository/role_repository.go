package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type RoleRepository interface {
	GetRoles(ctx context.Context, name string) ([]model.Role, error)
	GetRoleByID(ctx context.Context, id string) (*model.Role, error)
	CreateRole(ctx context.Context, role *model.Role) error
	UpdateRole(ctx context.Context, id string, updateMap map[string]any) error
	DeleteRole(ctx context.Context, id string) error
}

// struct implementasi
type roleRepository struct {
	db *gorm.DB
}

// constructor
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

// struct method
func (r *roleRepository) GetRoles(ctx context.Context, name string) ([]model.Role, error) {
	var roles []model.Role
	// query search sementara tanpa name
	query := r.db.WithContext(ctx).Preload("Tenant")

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	err := query.Find(&roles).Error

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) GetRoleByID(ctx context.Context, id string) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).Preload("Tenant").First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) CreateRole(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) UpdateRole(ctx context.Context, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.Role{}).Where("id = ?", id).Updates(updateMap).Error
}

func (r *roleRepository) DeleteRole(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Role{}).Error
}
