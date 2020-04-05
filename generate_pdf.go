package main

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"strings"
)

func GeneratePDF(html string) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.Dpi.Set(1200)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
