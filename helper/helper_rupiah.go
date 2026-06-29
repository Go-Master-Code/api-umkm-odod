package helper

import (
	"strings"

	"github.com/dustin/go-humanize"
)

func FormatRupiah(value float64) string {
	result := humanize.Commaf(value)
	return strings.ReplaceAll(result, ",", ".") // ganti , dengan . sebagai pemisah ribuan
}
