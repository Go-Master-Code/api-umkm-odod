package pdf

import (
	"bytes"
	"fmt"
	"strconv"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"github.com/go-pdf/fpdf"
)

func GenerateSalesReport(sales []model.Sale, query dto.SaleReportQuery, summary *dto.SalesReportSummary) (*bytes.Buffer, error) {
	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.AddPage() // WAJIB

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 7, "SALES REPORT", "", 1, "", false, 0, "")

	pdf.Ln(2) // beri jarak 24mm

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0, 4, fmt.Sprintf("Periode: %s s.d %s", query.StartDate, query.EndDate), "", 1, "", false, 0, "")
	pdf.Ln(2) // beri jarak 2mm

	// garis pemisah
	pdf.Ln(2)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)

	// local helper agar elemen summary sejajar
	writeSummaryRow := func(label string, value string) {
		pdf.CellFormat(35, 7, label, "", 0, "L", false, 0, "") // kasih space 50 untuk labelnya
		pdf.CellFormat(5, 7, ":", "", 0, "C", false, 0, "")
		pdf.CellFormat(5, 7, "Rp", "", 0, "C", false, 0, "")
		pdf.CellFormat(25, 7, value, "", 1, "R", false, 0, "")
	}

	pdf.Ln(2) // beri jarak 4mm
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 7, "SUMARRY", "", 1, "", false, 0, "")
	pdf.Ln(2) // beri jarak 4mm

	// total transaction dibuat tanpa currency Rp
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(35, 7, "Total Transaction", "", 0, "L", false, 0, "") // kasih space 50 untuk labelnya
	pdf.CellFormat(4, 7, ":", "", 0, "C", false, 0, "")
	pdf.CellFormat(5, 7, strconv.Itoa(int(summary.TotalTransaction)), "", 1, "L", false, 0, "")

	writeSummaryRow("Total Sales", helper.FormatRupiah(summary.TotalSales))
	writeSummaryRow("Total Discount", helper.FormatRupiah(summary.TotalDiscount))
	writeSummaryRow("Total Tax", helper.FormatRupiah(summary.TotalTax))
	writeSummaryRow("Grand Total", helper.FormatRupiah(summary.GrandTotal))
	pdf.Ln(2) // beri jarak 4mm

	// garis pemisah
	pdf.Ln(2)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)

	pdf.Ln(2) // beri jarak 2mm

	// row sales
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 6, "DETAIL TRANSACTION", "", 1, "", false, 0, "")
	pdf.Ln(4) // beri jarak 4mm

	// =====================================================
	// TABLE HEADER
	// =====================================================

	pdf.SetFont("Arial", "B", 10) // font style BOLD untuk header table

	headers := []struct {
		Title string
		Width float64
	}{
		{"Invoice", 30},
		{"Date", 25},
		{"Customer", 25},
		{"Cashier", 25},
		{"Subtotal", 22},
		{"Discount", 20},
		{"Tax", 19},
		{"Grand Total", 24},
	}

	for _, h := range headers {
		pdf.CellFormat(h.Width, 8, h.Title, "1", 0, "C", false, 0, "")
	}

	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10) // style biasa tanpa bold untuk isi tabel

	// isikan data row sales
	for _, row := range sales {
		pdf.CellFormat(30, 8, row.InvoiceNumber, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, row.CreatedAt.Format("02 Jan 2006"), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, row.CustomerName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, row.Cashier.FullName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(22, 8, helper.FormatRupiah(row.Subtotal), "1", 0, "R", false, 0, "")
		pdf.CellFormat(20, 8, helper.FormatRupiah(row.DiscountAmount), "1", 0, "R", false, 0, "")
		pdf.CellFormat(19, 8, helper.FormatRupiah(row.TaxAmount), "1", 0, "R", false, 0, "")
		pdf.CellFormat(24, 8, helper.FormatRupiah(row.GrandTotal), "1", 0, "R", false, 0, "")
		pdf.Ln(-1) // line break seperti enter
	}

	var buf bytes.Buffer

	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
