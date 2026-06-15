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
type ActivityLogService interface {
	GetAllActivityLog(ctx context.Context, query dto.GetAllActivityLogResponseQuery) ([]dto.ActivityLogResponse, int64, error)
	CreateActivityLog(ctx context.Context, module string, action string, description string, referenceID string, referenceNumber string) error
}

// struct implementasi
type activityLogService struct {
	repo repository.ActivityLogRepository
}

// constructor
func NewActivityLogService(repo repository.ActivityLogRepository) ActivityLogService {
	return &activityLogService{
		repo: repo,
	}
}

// struct method
func (s *activityLogService) GetAllActivityLog(ctx context.Context, query dto.GetAllActivityLogResponseQuery) ([]dto.ActivityLogResponse, int64, error) {
	// ambil tenantID dari jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	logs, total, err := s.repo.GetAllActivityLog(ctx, tenantID, query)
	if err != nil {
		return nil, 0, err
	}

	// convert model to dto
	logsDTO := helper.ConvertToDTOActivityLogPlural(logs)

	return logsDTO, total, nil
}

func (s *activityLogService) CreateActivityLog(ctx context.Context, module string, action string, description string, referenceID string, referenceNumber string) error {
	// ambil tenantID dan userID dari jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	userID := ctx.Value(constants.ContextUserID).(string)

	activityLog := model.ActivityLog{
		ID:              uuid.NewString(),
		TenantID:        tenantID,
		UserID:          userID,
		Module:          module,
		Action:          action,
		Description:     description,
		ReferenceID:     referenceID,
		ReferenceNumber: referenceNumber,
	}

	err := s.repo.CreateActivityLog(ctx, &activityLog)
	if err != nil {
		return err
	}

	return nil
}
