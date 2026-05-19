package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type TenantRepository interface {
	GetTenants(ctx context.Context, name string) ([]model.Tenant, error)
	GetTenantByID(ctx context.Context, id string) (*model.Tenant, error)
	CreateTenant(ctx context.Context, tenant *model.Tenant) error
	UpdateTenant(ctx context.Context, id string, updateMap map[string]any) error
	DeleteTenant(ctx context.Context, id string) error
}

// struct implementasi
type tenantRepository struct {
	db *gorm.DB
}

// constructor
func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{
		db: db,
	}
}

// struct method
func (r *tenantRepository) GetTenants(ctx context.Context, name string) ([]model.Tenant, error) {
	var tenants []model.Tenant
	// query search sementara (tanpa parameter name dll)
	query := r.db.WithContext(ctx)

	if name != "" { // jika param nama tidak kosong, tambahkan ke query
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	err := query.Find(&tenants).Error

	if err != nil {
		return nil, err // default value dari slice model = nil
	}
	return tenants, nil
}

func (r *tenantRepository) GetTenantByID(ctx context.Context, id string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.WithContext(ctx).First(&tenant, "id = ?", id).Error // lebih idiomatik daripada pake klausa where
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) CreateTenant(ctx context.Context, tenant *model.Tenant) error {
	return r.db.WithContext(ctx).Create(tenant).Error
}

func (r *tenantRepository) UpdateTenant(ctx context.Context, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.Tenant{}).Where("id = ?", id).Updates(updateMap).Error
}

func (r *tenantRepository) DeleteTenant(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Tenant{}).Error
}
