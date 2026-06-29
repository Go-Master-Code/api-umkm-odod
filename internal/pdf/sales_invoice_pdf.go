package pdf

import (
	"bytes"
	"fmt"
	"umkm-odod/helper"
	"umkm-odod/internal/model"

	"github.com/go-pdf/fpdf"
)

func GenerateSalesInvoice(sale model.Sale) (*bytes.Buffer, error) {
	pdf := fpdf.New("P", "mm", "", "")

	pdf.AddPageFormat(
		"P",
		fpdf.SizeType{
			Wd: 80,  // lebar 80 mm
			Ht: 200, // walau height 200 mm, printer hanya akan print bagian kertas yang ada isi teksnya
		},
	)

	// helper lokal
	writeAmountRow := func(label string, amount float64) {
		pdf.CellFormat(40, 5, label, "", 0, "L", false, 0, "")                      // print dari sisi kiri kertas
		pdf.CellFormat(0, 5, helper.FormatRupiah(amount), "", 1, "R", false, 0, "") // print dari sisi kanan kertas
	}

	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(0, 6, sale.Tenant.Name, "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(0, 4, sale.Tenant.Address, "", 1, "C", false, 0, "")

	// garis pemisah
	pdf.Ln(2)

	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")

	pdf.Ln(2)

	// header sales
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(0, 4, fmt.Sprintf("Invoice: %s", sale.InvoiceNumber), "", 1, "", false, 0, "")
	pdf.CellFormat(0, 4, fmt.Sprintf("Date: %s", sale.CreatedAt.Format("02 Jan 2006 15:04")), "", 1, "", false, 0, "")
	pdf.CellFormat(0, 4, fmt.Sprintf("Cashier: %s", sale.Cashier.FullName), "", 1, "", false, 0, "")
	pdf.CellFormat(0, 4, fmt.Sprintf("Customer: %s", sale.CustomerName), "", 1, "", false, 0, "")

	// garis pemisah
	pdf.Ln(1)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)

	// detil sale item
	pdf.SetFont("Arial", "", 8)
	for _, item := range sale.SaleItems {
		itemName := item.ItemNameSnapshot

		if item.VariantNameSnapshot != "" { // format item per row: itemName - itemVariant
			itemName += " - " + item.VariantNameSnapshot
		}

		// Nama barang
		pdf.MultiCell(0, 4, itemName, "", "L", false)

		// Qty x Harga
		leftText := fmt.Sprintf(
			"%.0f x %s", // f artinya float, s = string
			item.Qty,
			helper.FormatRupiah(item.UnitPrice),
		)

		// Subtotal per baris data
		rightText := helper.FormatRupiah(item.Subtotal)

		// cetak left text (qty * harga) di sebelah kiri, subtotal di sebelah kanan
		pdf.CellFormat(40, 4, leftText, "", 0, "L", false, 0, "")
		pdf.CellFormat(0, 4, rightText, "", 1, "R", false, 0, "")

		pdf.Ln(1)
	}

	// garis pemisah
	pdf.Ln(1)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)

	// summary
	writeAmountRow("Subtotal: ", sale.Subtotal)
	writeAmountRow("Discount: ", sale.DiscountAmount)
	writeAmountRow("Tax: ", sale.TaxAmount)

	// grand total
	pdf.SetFont("Arial", "B", 10)
	writeAmountRow("Grand Total: ", sale.GrandTotal)

	// garis pemisah
	pdf.Ln(1)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)

	// footer tenant
	pdf.SetFont("Arial", "", 8)
	// untuk data berupa text dari db, pakai multicell agar go bisa mendeteksi space antar baris
	pdf.MultiCell(0, 4, sale.Tenant.ReceiptFooter, "", "C", false)
	pdf.Ln(2)                                                                                  //space
	pdf.MultiCell(0, 4, fmt.Sprintf("Kritik & Saran:\n%s", sale.Tenant.Phone), "", "C", false) // untuk \n harus digunakan dengan pdf.MultiCell agar bisa ada space antar baris

	var buf bytes.Buffer

	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
