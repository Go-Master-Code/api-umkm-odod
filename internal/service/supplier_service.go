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
type SupplierService interface {
	GetSuppliers(ctx context.Context, name string) ([]dto.SupplierResponse, error)
	GetSupplierByID(ctx context.Context, id string) (dto.SupplierResponse, error)
	CreateSupplier(ctx context.Context, req dto.CreateSupplierRequest) (dto.SupplierResponse, error)
	UpdateSupplier(ctx context.Context, id string, req dto.UpdateSupplierRequest) (dto.SupplierResponse, error)
	DeleteSupplier(ctx context.Context, id string) (dto.SupplierResponse, error)
}

// struct implementasi
type supplierService struct {
	repo repository.SupplierRepository
}

// constructor
func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierService{
		repo: repo,
	}
}

// struct implementasi
func (s *supplierService) GetSuppliers(ctx context.Context, name string) ([]dto.SupplierResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	suppliers, err := s.repo.GetSuppliers(ctx, tenantID, name)
	if err != nil {
		return nil, err
	}

	// convert to dto supplier
	suppliersDTO := helper.ConvertToDTOSupplierPlural(suppliers)
	return suppliersDTO, nil
}

func (s *supplierService) GetSupplierByID(ctx context.Context, id string) (dto.SupplierResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	supplier, err := s.repo.GetSupplierByID(ctx, tenantID, id)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// convert model to dto
	supplierDTO := helper.ConvertToDTOSupplierSingle(supplier)
	return supplierDTO, nil
}

func (s *supplierService) CreateSupplier(ctx context.Context, req dto.CreateSupplierRequest) (dto.SupplierResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// convert req to model for param
	supplier := model.Supplier{
		ID:       uuid.NewString(),
		TenantID: tenantID,
		Name:     req.Name,
		Address:  req.Address,
		Phone:    req.Phone,
	}

	err := s.repo.CreateSupplier(ctx, &supplier)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// get data untuk preload relasi
	newSupplier, err := s.repo.GetSupplierByID(ctx, tenantID, supplier.ID)

	if err != nil {
		return dto.SupplierResponse{}, err
	}

	supplierDTO := helper.ConvertToDTOSupplierSingle(newSupplier)

	return supplierDTO, nil
}

func (s *supplierService) UpdateSupplier(ctx context.Context, id string, req dto.UpdateSupplierRequest) (dto.SupplierResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// map update data di repo
	var updateMap = map[string]any{}

	// cek perubahan data pada request body
	if req.Address != nil {
		updateMap["address"] = req.Address
	}
	if req.Name != nil {
		updateMap["name"] = req.Name
	}
	if req.Phone != nil {
		updateMap["phone"] = req.Phone
	}

	err := s.repo.UpdateSupplier(ctx, tenantID, id, updateMap)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// get udpated data from db
	updateSupplier, err := s.repo.GetSupplierByID(ctx, tenantID, id)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// convert data to dto
	supplierDTO := helper.ConvertToDTOSupplierSingle(updateSupplier)

	return supplierDTO, nil
}

func (s *supplierService) DeleteSupplier(ctx context.Context, id string) (dto.SupplierResponse, error) {
	// get tenantID from context
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data dulu untuk ditampilkan setelah data berhasil di delete
	supplier, err := s.repo.GetSupplierByID(ctx, tenantID, id)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// delete data
	err = s.repo.DeleteSupplier(ctx, tenantID, id)

	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// jika berhasil, convert model to dto
	supplierDTO := helper.ConvertToDTOSupplierSingle(supplier)
	return supplierDTO, nil
}
