package service

import (
	"context"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/repository"

	"github.com/google/uuid"
)

// interface
type RoleService interface {
	GetRoles(ctx context.Context, name string) ([]dto.RoleResponse, error)
	GetRoleByID(ctx context.Context, id string) (dto.RoleResponse, error)
	CreateRole(ctx context.Context, req dto.CreateRoleRequest) (dto.RoleResponse, error)
	UpdateRole(ctx context.Context, id string, req dto.UpdateRoleRequest) (dto.RoleResponse, error)
	DeleteRole(ctx context.Context, id string) (dto.RoleResponse, error)
}

// struct impelemtasi
type roleService struct {
	repo repository.RoleRepository
}

// constructor
func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{
		repo: repo,
	}
}

// struct method
func (s *roleService) GetRoles(ctx context.Context, name string) ([]dto.RoleResponse, error) {
	roles, err := s.repo.GetRoles(ctx, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	rolesDTO := helper.ConvertToDTORolePlural(roles)
	return rolesDTO, nil
}

func (s *roleService) GetRoleByID(ctx context.Context, id string) (dto.RoleResponse, error) {
	role, err := s.repo.GetRoleByID(ctx, id)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(role)
	return roleDTO, nil
}

func (s *roleService) CreateRole(ctx context.Context, req dto.CreateRoleRequest) (dto.RoleResponse, error) {
	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// payload sementara
	// tenantID := "f27e441f-5385-4b8d-b2e2-88b8615a4634"

	// convert dto to model
	role := model.Role{
		TenantID: tenantID,
		ID:       uuid.NewString(),
		Name:     req.Name,
	}

	err := s.repo.CreateRole(ctx, &role)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// get role by id (preload tenant) agar nama tenant muncul
	newRole, err := s.repo.GetRoleByID(ctx, role.ID)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(newRole)
	return roleDTO, nil
}

func (s *roleService) UpdateRole(ctx context.Context, id string, req dto.UpdateRoleRequest) (dto.RoleResponse, error) {
	// mapping update data ke map
	var updateMap = map[string]any{}

	if req.Name != nil {
		updateMap["name"] = req.Name
	}

	// update data
	err := s.repo.UpdateRole(ctx, id, updateMap)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// get data yang sudah diupdate untuk return value
	updateRole, err := s.repo.GetRoleByID(ctx, id)

	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(updateRole)
	return roleDTO, nil
}

func (s *roleService) DeleteRole(ctx context.Context, id string) (dto.RoleResponse, error) {
	// get data dulu sebelum di delete
	role, err := s.repo.GetRoleByID(ctx, id)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// delete data
	err = s.repo.DeleteRole(ctx, id)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(role)
	return roleDTO, nil
}
