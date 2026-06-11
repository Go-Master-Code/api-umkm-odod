package service

import (
	"context"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"
)

// interface
type ReportService interface {
	GetSalesReport(ctx context.Context, query dto.SaleReportQuery) (dto.SalesReportResponse, error)
}

// struct implementasi
type reportService struct {
	repo repository.ReportRepository
}

// constructor
func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{
		repo: repo,
	}
}

func (s *reportService) GetSalesReport(ctx context.Context, query dto.SaleReportQuery) (dto.SalesReportResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	reportSummary, err := s.repo.GetSalesReportSummary(ctx, tenantID, query.StartDate, query.EndDate)
	if err != nil {
		return dto.SalesReportResponse{}, err
	}

	salesReport, err := s.repo.GetSalesReport(ctx, tenantID, query.StartDate, query.EndDate)
	if err != nil {
		return dto.SalesReportResponse{}, err
	}

	// convert salesReport to dto
	salesReportDTO := helper.ConvertToDTOSalePlural(salesReport)

	// masukkan data summary dan transaction ke sales report response
	salesReportFull := dto.SalesReportResponse{
		Summary:      *reportSummary,
		Transactions: salesReportDTO,
	}

	return salesReportFull, nil
}
