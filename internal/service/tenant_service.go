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
type TenantService interface {
	GetTenants(ctx context.Context, name string) ([]dto.TenantResponse, error)
	GetTenantByID(ctx context.Context, id string) (dto.TenantResponse, error)
	CreateTenant(ctx context.Context, req dto.CreateTenantRequest) (dto.TenantResponse, error)
	UpdateTenant(ctx context.Context, id string, req dto.UpdateTenantRequest) (dto.TenantResponse, error)
	DeleteTenant(ctx context.Context, id string) (dto.TenantResponse, error)
}

// struct implementasi
type tenantService struct {
	repo repository.TenantRepository
}

// constructor
func NewTenantService(repo repository.TenantRepository) TenantService {
	return &tenantService{
		repo: repo,
	}
}

// struct method
func (s *tenantService) GetTenants(ctx context.Context, name string) ([]dto.TenantResponse, error) {
	tenants, err := s.repo.GetTenants(ctx, name)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	tenantsDTO := helper.ConvertToDTOTenantPlural(tenants)
	return tenantsDTO, nil
}

func (s *tenantService) GetTenantByID(ctx context.Context, id string) (dto.TenantResponse, error) {
	tenant, err := s.repo.GetTenantByID(ctx, id)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	// convert model to dto
	tenantDTO := helper.ConvertToDTOTenantSingle(tenant)
	return tenantDTO, nil
}

func (s *tenantService) CreateTenant(ctx context.Context, req dto.CreateTenantRequest) (dto.TenantResponse, error) {
	// parsing request body ke model
	tenant := model.Tenant{
		ID:      uuid.NewString(), // generate uuid pakai lib
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}

	err := s.repo.CreateTenant(ctx, &tenant)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	// convert model to dto as a return value
	tenantDTO := helper.ConvertToDTOTenantSingle(&tenant)
	return tenantDTO, nil
}

func (s *tenantService) UpdateTenant(ctx context.Context, id string, req dto.UpdateTenantRequest) (dto.TenantResponse, error) {
	// buat map untuk param func repo
	var updateMap = map[string]any{}

	// cek setiap atribut model apakah ada isinya, jika ada isinya, update field di repo
	if req.Name != nil {
		updateMap["name"] = *req.Name
	}
	if req.Phone != nil {
		updateMap["phone"] = *req.Phone
	}
	if req.Address != nil {
		updateMap["address"] = *req.Address
	}

	err := s.repo.UpdateTenant(ctx, id, updateMap)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	// ambil ulang data yang sudah diupdate
	updateTenant, err := s.repo.GetTenantByID(ctx, id)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	// convert model -> dto sebagai response
	tenantDTO := helper.ConvertToDTOTenantSingle(updateTenant)
	return tenantDTO, nil
}

func (s *tenantService) DeleteTenant(ctx context.Context, id string) (dto.TenantResponse, error) {
	// ambil data yang ingin di delete
	tenant, err := s.repo.GetTenantByID(ctx, id)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	err = s.repo.DeleteTenant(ctx, id)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	// convert model to dto
	tenantDTO := helper.ConvertToDTOTenantSingle(tenant)

	return tenantDTO, nil
}
