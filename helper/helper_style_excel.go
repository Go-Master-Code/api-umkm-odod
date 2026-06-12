package helper

import "github.com/xuri/excelize/v2"

func BoldStyle(f *excelize.File) int {
	// bold style
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	if err != nil {
		return 0
	}

	return boldStyle

}

func HeaderStyle(f *excelize.File) int {
	// header style : bold + center
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})

	if err != nil {
		return 0
	}

	return headerStyle
}

func CurrencyStyle(f *excelize.File) int {
	// formatting rupiah tapi angka tetap bertipe number agar tetap bisa diagregat di excel nanti
	format := `"Rp" #,##0`
	currencyStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &format,
	})

	if err != nil {
		return 0
	}

	return currencyStyle
}
