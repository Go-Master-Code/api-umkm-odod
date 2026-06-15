package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ActivityLogRepository interface {
	GetAllActivityLog(ctx context.Context, tenantID string, query dto.GetAllActivityLogResponseQuery) ([]model.ActivityLog, int64, error)
	CreateActivityLog(ctx context.Context, log *model.ActivityLog) error
}

// struct implementasi
type activityLogRepository struct {
	db *gorm.DB
}

// constructor
func NewActivityLogRepository(db *gorm.DB) ActivityLogRepository {
	return &activityLogRepository{
		db: db,
	}
}

// struct method
func (r *activityLogRepository) CreateActivityLog(ctx context.Context, log *model.ActivityLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *activityLogRepository) GetAllActivityLog(ctx context.Context, tenantID string, query dto.GetAllActivityLogResponseQuery) ([]model.ActivityLog, int64, error) {
	var activityLogs []model.ActivityLog
	var total int64

	offset := (query.Page - 1) * query.Limit

	baseQuery := r.db.WithContext(ctx).Model(&model.ActivityLog{}).Where("tenant_id = ?", tenantID)

	//search
	if query.Search != "" {
		search := "%" + query.Search + "%"
		baseQuery = baseQuery.Where(`module LIKE ? OR action LIKE ? OR description LIKE ? OR reference_number LIKE ?`, search, search, search, search)
	}

	// count total data
	err := baseQuery.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// get data
	err = baseQuery.Preload("Tenant").Preload("User").
		Order("created_at DESC").
		Limit(query.Limit).
		Offset(offset).
		Find(&activityLogs).Error

	if err != nil {
		return nil, 0, err
	}

	// jika sukses
	return activityLogs, total, nil
}
