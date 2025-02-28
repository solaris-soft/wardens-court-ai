package pdfReader

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

// ExtractText extracts text content from a PDF file between specified pages.
// If startPage and endPage are both 0, extracts from the entire document.
func ExtractText(pdfPath string, startPage int, endPage int) (string, error) {
	f, r, err := pdf.Open(pdfPath)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil

}
