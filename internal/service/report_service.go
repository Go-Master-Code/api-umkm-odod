package service

import (
	"context"
	"fmt"
	"time"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"

	"github.com/xuri/excelize/v2"
)

// interface
type ReportService interface {
	GetSalesReport(ctx context.Context, query dto.SaleReportQuery) (dto.SalesReportResponse, error)
	GetPurchaseReport(ctx context.Context, query dto.PurchaseReportQuery) (dto.PurchaseReportResponse, error)
	GetStockReport(ctx context.Context, query dto.StockReportQuery) ([]dto.StockReportResponse, int64, error)
	// export report to xlsx
	ExportSalesReport(ctx context.Context, query dto.SaleReportQuery) (*excelize.File, error)
	ExportPurchaseReport(ctx context.Context, query dto.PurchaseReportQuery) (*excelize.File, error)
	ExportStockReport(ctx context.Context, query dto.StockReportQuery) (*excelize.File, error)
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

func (s *reportService) GetStockReport(ctx context.Context, query dto.StockReportQuery) ([]dto.StockReportResponse, int64, error) {
	// get tenantID by context
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	result, total, err := s.repo.GetStockReport(ctx, tenantID, query)

	if err != nil {
		return nil, 0, err
	}

	// looping untuk menentukan status barang
	for i := range result {
		if result[i].CurrentStock < result[i].MinimumStock {
			result[i].Status = "LOW"
		} else {
			result[i].Status = "NORMAL"
		}
	}

	return result, total, nil
}

// ====================================================
// ===========SECTION EXPORT REPORT TO EXCEL===========
// ====================================================
func (s *reportService) ExportSalesReport(ctx context.Context, query dto.SaleReportQuery) (*excelize.File, error) {
	// get tenantID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data sales dulu, []model.Sale
	sales, err := s.repo.GetSalesReport(ctx, tenantID, query.StartDate, query.EndDate)

	if err != nil {
		return nil, err
	}

	// buat excel workbook
	f := excelize.NewFile()

	sheetName := "Sales Report"

	f.SetSheetName("Sheet1", sheetName) // ganti nama sheet

	// ambil data summary dari sales repo
	summary, err := s.repo.GetSalesReportSummary(ctx, tenantID, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	// SUMMARY
	f.SetCellValue(sheetName, "A1", "SALES REPORT")
	f.SetCellValue(sheetName, "A2", fmt.Sprintf("Periode: %s s/d %s", query.StartDate, query.EndDate))
	f.SetCellValue(sheetName, "A4", "SUMMARY")
	f.SetCellValue(sheetName, "A5", "Total Transaction")
	f.SetCellValue(sheetName, "B5", summary.TotalTransaction)
	f.SetCellValue(sheetName, "A6", "Total Sales")
	f.SetCellValue(sheetName, "B6", summary.TotalSales)
	f.SetCellValue(sheetName, "A7", "Total Discount")
	f.SetCellValue(sheetName, "B7", summary.TotalDiscount)
	f.SetCellValue(sheetName, "A8", "Total Tax")
	f.SetCellValue(sheetName, "B8", summary.TotalTax)
	f.SetCellValue(sheetName, "A9", "Grand Total")
	f.SetCellValue(sheetName, "B9", summary.GrandTotal)
	f.SetCellValue(sheetName, "A11", "SALES DETAILS")

	// styling bold
	f.SetCellStyle(sheetName, "A1", "A1", helper.BoldStyle(f))
	f.SetCellStyle(sheetName, "A4", "A4", helper.BoldStyle(f))
	f.SetCellStyle(sheetName, "A11", "A11", helper.BoldStyle(f))

	// styling currency untuk nominal di summary
	f.SetCellStyle(sheetName, "B6", "B9", helper.CurrencyStyle(f))

	// header report
	headers := []string{
		"Invoice Number",
		"Date",
		"Customer",
		"Cashier",
		"Subtotal",
		"Discount",
		"Tax",
		"Grand Total",
		"Payment Method",
		"Status",
	}

	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 12) // mulai dari row 12 karena bagian atas sheet dipakai summary
		f.SetCellValue(sheetName, cell, header)
	}

	// style bold untuk header dimulai dari row 12
	f.SetCellStyle(sheetName, "A12", "J12", helper.HeaderStyle(f))

	// atur width tiap kolom: parameter:sheetname,startingColumn,endColumn,width
	f.SetColWidth(sheetName, "A", "A", 17)
	f.SetColWidth(sheetName, "B", "B", 17)
	f.SetColWidth(sheetName, "C", "C", 20)
	f.SetColWidth(sheetName, "D", "D", 20)
	f.SetColWidth(sheetName, "E", "H", 15)
	f.SetColWidth(sheetName, "I", "I", 17)
	f.SetColWidth(sheetName, "J", "J", 10)

	row := 13 // dimulai dari row 13 karena di atasnya ada header + summary

	// looping data sales -> isikan ke tiap row excel
	for _, sale := range sales {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), sale.InvoiceNumber) // %d artinya data (var row) bernilai int
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), sale.CreatedAt)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), sale.CustomerName)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), sale.Cashier.FullName)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), sale.Subtotal)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), sale.DiscountAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), sale.TaxAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), sale.GrandTotal)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), sale.PaymentMethod)
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), sale.PaymentStatus)

		row++ // increment to the next row
	}

	// ambil last row untuk apply style rupiah pada number
	// lastRow := row - 1 // row terakhir yang berisi data

	// label total untuk agregat data
	labelTotal := fmt.Sprintf("D%d", row)

	// caption total untuk agregasi data subtotal,discount,tax,dan grandtotal
	f.SetCellValue(sheetName, labelTotal, "Total")

	// styling header untuk agregat total
	f.SetCellStyle(sheetName, labelTotal, labelTotal, helper.HeaderStyle(f))

	// isikan data summary pada row paling bawah (row kosong setelah semua transaksi tertulis)
	f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), summary.TotalSales)
	f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), summary.TotalDiscount)
	f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), summary.TotalTax)
	f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), summary.GrandTotal)
	f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), fmt.Sprintf("(%d transaction)", summary.TotalTransaction))

	// styling header untuk total transaksi
	f.SetCellStyle(sheetName, fmt.Sprintf("I%d", row), fmt.Sprintf("I%d", row), helper.HeaderStyle(f))

	// styling currency untuk nominal
	f.SetCellStyle(sheetName, "E13", fmt.Sprintf("H%d", row), helper.CurrencyStyle(f))

	// 	Artinya mulai row 13 (data sales pertama)

	// E = Subtotal
	// F = Discount
	// G = Tax
	// H = Grand Total

	// semua otomatis format rupiah.

	// return excel file
	return f, nil
}

func (s *reportService) ExportPurchaseReport(ctx context.Context, query dto.PurchaseReportQuery) (*excelize.File, error) {
	// get tenantID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data purchase dulu, []model.Purchase
	purchases, err := s.repo.GetPurchaseReport(ctx, tenantID, query.StartDate, query.EndDate)

	if err != nil {
		return nil, err
	}

	// buat excel workbook
	f := excelize.NewFile()

	sheetName := "Purchase Report"

	f.SetSheetName("Sheet1", sheetName) // ganti nama sheet

	// ambil data summary dari purchases repo
	summary, err := s.repo.GetPurchaseReportSummary(ctx, tenantID, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	// SUMMARY
	f.SetCellValue(sheetName, "A1", "PURCHASE REPORT")
	f.SetCellValue(sheetName, "A2", fmt.Sprintf("Periode: %s s/d %s", query.StartDate, query.EndDate))
	f.SetCellValue(sheetName, "A4", "SUMMARY")
	f.SetCellValue(sheetName, "A5", "Total Transaction")
	f.SetCellValue(sheetName, "B5", summary.TotalTransaction)
	f.SetCellValue(sheetName, "A6", "Total Sales")
	f.SetCellValue(sheetName, "B6", summary.TotalPurchase)
	f.SetCellValue(sheetName, "A7", "Total Discount")
	f.SetCellValue(sheetName, "B7", summary.TotalDiscount)
	f.SetCellValue(sheetName, "A8", "Total Tax")
	f.SetCellValue(sheetName, "B8", summary.TotalTax)
	f.SetCellValue(sheetName, "A9", "Grand Total")
	f.SetCellValue(sheetName, "B9", summary.GrandTotal)
	f.SetCellValue(sheetName, "A11", "PURCHASE DETAILS")

	// styling bold
	f.SetCellStyle(sheetName, "A1", "A1", helper.BoldStyle(f))
	f.SetCellStyle(sheetName, "A4", "A4", helper.BoldStyle(f))
	f.SetCellStyle(sheetName, "A11", "A11", helper.BoldStyle(f))

	// styling currency untuk nominal di summary
	f.SetCellStyle(sheetName, "B6", "B9", helper.CurrencyStyle(f))

	// header report
	headers := []string{
		"Purchase No.",
		"Date",
		"Supplier",
		"Invoice No.",
		"Creator",
		"Subtotal",
		"Discount",
		"Tax",
		"Grand Total",
		"Notes",
	}

	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 12) // mulai dari row 12 karena bagian atas sheet dipakai summary
		f.SetCellValue(sheetName, cell, header)
	}

	// style bold untuk header dimulai dari row 12
	f.SetCellStyle(sheetName, "A12", "J12", helper.HeaderStyle(f))

	// atur width tiap kolom: parameter:sheetname,startingColumn,endColumn,width
	f.SetColWidth(sheetName, "A", "A", 17)
	f.SetColWidth(sheetName, "B", "B", 17)
	f.SetColWidth(sheetName, "C", "C", 20)
	f.SetColWidth(sheetName, "D", "D", 17)
	f.SetColWidth(sheetName, "E", "E", 20)
	f.SetColWidth(sheetName, "F", "I", 15)
	f.SetColWidth(sheetName, "J", "J", 20)

	row := 13 // dimulai dari row 13 karena di atasnya ada header + summary

	// looping data purchases -> isikan ke tiap row excel
	for _, p := range purchases {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), p.PurchaseNumber) // %d artinya data (var row) bernilai int
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), p.CreatedAt)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), p.Supplier.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), p.InvoiceNumber)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), p.Creator.FullName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), p.Subtotal)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), p.DiscountAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), p.TaxAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), p.GrandTotal)
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), p.Notes)

		row++ // increment to the next row
	}

	// ambil last row untuk apply style rupiah pada number
	// lastRow := row - 1 // row terakhir yang berisi data

	// label total untuk agregat data
	labelTotal := fmt.Sprintf("E%d", row)

	// caption total untuk agregasi data subtotal,discount,tax,dan grandtotal
	f.SetCellValue(sheetName, labelTotal, "Total")

	// styling header untuk agregat total
	f.SetCellStyle(sheetName, labelTotal, labelTotal, helper.HeaderStyle(f))

	// isikan data summary pada row paling bawah (row kosong setelah semua transaksi tertulis)
	f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), summary.TotalPurchase)
	f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), summary.TotalDiscount)
	f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), summary.TotalTax)
	f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), summary.GrandTotal)
	f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), fmt.Sprintf("(%d transaction)", summary.TotalTransaction))

	// styling header untuk total transaksi
	f.SetCellStyle(sheetName, fmt.Sprintf("J%d", row), fmt.Sprintf("J%d", row), helper.HeaderStyle(f))

	// styling currency untuk nominal
	f.SetCellStyle(sheetName, "F13", fmt.Sprintf("I%d", row), helper.CurrencyStyle(f))

	return f, nil
}

func (s *reportService) ExportStockReport(ctx context.Context, query dto.StockReportQuery) (*excelize.File, error) {
	// tenantID dari jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// ambil data stock per barang
	stock, _, err := s.repo.GetStockReport(ctx, tenantID, query)
	if err != nil {
		return nil, err
	}

	// summary report stock
	summary, err := s.repo.GetStockReportSummary(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	// buat excel workbook
	f := excelize.NewFile()

	sheetName := "Stock Report"

	f.SetSheetName("Sheet1", sheetName) // ganti nama sheet

	// generated at
	now := time.Now()
	dateStr := now.Format("2006-01-02")

	// SUMMARY
	f.SetCellValue(sheetName, "A1", "STOCK REPORT")
	f.SetCellValue(sheetName, "A2", "Generated At")
	f.SetCellValue(sheetName, "B2", dateStr)
	f.SetCellValue(sheetName, "A3", "Total Variants")
	f.SetCellValue(sheetName, "B3", summary.TotalVariants)
	f.SetCellValue(sheetName, "A4", "Low Stock Items")
	f.SetCellValue(sheetName, "B4", summary.LowStockItems)
	f.SetCellValue(sheetName, "A6", "STOCK DETAILS")

	// styling bold
	f.SetCellStyle(sheetName, "A1", "A1", helper.BoldStyle(f))
	f.SetCellStyle(sheetName, "A6", "A6", helper.BoldStyle(f))

	// header report
	headers := []string{
		"Item Name",
		"Variant",
		"SKU",
		"Stock",
		"Min. Stock",
		"Status",
	}

	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 7) // mulai dari row 7 karena bagian atas sheet dipakai summary
		f.SetCellValue(sheetName, cell, header)
	}

	// style bold untuk header dimulai dari row 7
	f.SetCellStyle(sheetName, "A7", "J7", helper.HeaderStyle(f))

	// atur width tiap kolom: parameter:sheetname,startingColumn,endColumn,width
	f.SetColWidth(sheetName, "A", "A", 30)
	f.SetColWidth(sheetName, "B", "B", 25)
	f.SetColWidth(sheetName, "C", "E", 15)
	f.SetColWidth(sheetName, "F", "F", 10)

	row := 7 // dimulai dari row 7 karena di atasnya ada header + summary

	// default status
	status := "NORMAL"

	// looping data purchases -> isikan ke tiap row excel
	for _, s := range stock {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), s.ItemName) // %d artinya data (var row) bernilai int
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), s.VariantName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), s.SKU)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), s.CurrentStock)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), s.MinimumStock)

		if s.CurrentStock < s.MinimumStock {
			status = "LOW"
		} else {
			status = "NORMAL"
		}
		s.Status = status

		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), s.Status)
		row++ // increment to the next row
	}

	return f, nil
}
