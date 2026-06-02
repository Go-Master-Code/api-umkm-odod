package repository

import (
	"umkm-odod/internal/model"

	"context"

	"gorm.io/gorm"
)

// interface
type SupplierRepository interface {
	GetSuppliers(ctx context.Context, tenantID string, name string) ([]model.Supplier, error)
	GetSupplierByID(ctx context.Context, tenantID string, id string) (*model.Supplier, error)
	CreateSupplier(ctx context.Context, supplier *model.Supplier) error
	UpdateSupplier(ctx context.Context, tenantID, id string, updateMap map[string]any) error
	DeleteSupplier(ctx context.Context, tenantID, id string) error
	IsSupplierNameExist(ctx context.Context, tenantID, name string) (bool, error)
}

// struct implementasi
type supplierRepository struct {
	db *gorm.DB
}

// constructor
func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{
		db: db,
	}
}

// struct method
func (r *supplierRepository) GetSuppliers(ctx context.Context, tenantID string, name string) ([]model.Supplier, error) {
	var suppliers []model.Supplier

	// query dasar
	query := r.db.WithContext(ctx).Preload("Tenant").Where("tenant_id = ?", tenantID)

	// jika ada parameter name
	if name != "" {
		// tambahkan klausa where pada query dasar
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// eksekusi query
	err := query.Find(&suppliers).Error
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *supplierRepository) GetSupplierByID(ctx context.Context, tenantID string, id string) (*model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.WithContext(ctx).Preload("Tenant").Where("tenant_id = ? and id = ?", tenantID, id).First(&supplier).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (r *supplierRepository) CreateSupplier(ctx context.Context, supplier *model.Supplier) error {
	return r.db.WithContext(ctx).Create(supplier).Error
}

func (r *supplierRepository) UpdateSupplier(ctx context.Context, tenantID, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.Supplier{}).Where("id = ? and tenant_id = ?", id, tenantID).Updates(updateMap).Error
}

func (r *supplierRepository) DeleteSupplier(ctx context.Context, tenantID, id string) error {
	return r.db.WithContext(ctx).Where("id = ? and tenant_id = ?", id, tenantID).Delete(&model.Supplier{}).Error
}

func (r *supplierRepository) IsSupplierNameExist(ctx context.Context, tenantID, name string) (bool, error) {
	var count int64
	// cek apakah nama suppllier sudah ada di suatu tenant
	err := r.db.WithContext(ctx).Where("tenant_id = ? and name = ?", tenantID, name).Count(&count).Error
	if err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}

	// jika name supplier tertentu belum ada di tenantID ini
	return false, nil
}
