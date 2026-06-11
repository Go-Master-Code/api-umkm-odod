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
	GetPurchaseReport(ctx context.Context, query dto.PurchaseReportQuery) (dto.PurchaseReportResponse, error)
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

func (s *reportService) GetPurchaseReport(ctx context.Context, query dto.PurchaseReportQuery) (dto.PurchaseReportResponse, error) {
	// get tenantID by context
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	reportSummary, err := s.repo.GetPurchaseReportSummary(ctx, tenantID, query.StartDate, query.EndDate)
	if err != nil {
		return dto.PurchaseReportResponse{}, err
	}

	// get data penjualan
	purchaseReport, err := s.repo.GetPurchaseReport(ctx, tenantID, query.StartDate, query.EndDate)
	if err != nil {
		return dto.PurchaseReportResponse{}, err
	}

	// convert purchase report to dto
	purchaseReportDTO := helper.ConvertToDTOPurchasePlural(purchaseReport)

	// masukkan data summary dan master-detil purchase ke dtoPurchaseReportResponse
	purchaseReportFull := dto.PurchaseReportResponse{
		Summary:      *reportSummary,
		Transactions: purchaseReportDTO,
	}

	return purchaseReportFull, nil
}
