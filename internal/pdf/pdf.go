package pdfReader

import (
	"bytes"
	"io"
	"os"

	"github.com/ledongthuc/pdf"
)

// ExtractText extracts text content from a PDF file between specified pages.
// If startPage and endPage are both 0, extracts from the entire document.
// If endPage is 0, extracts from startPage to the end of the document.
func ExtractText(pdfPath string, startPage int, endPage int) (string, error) {
	var buf bytes.Buffer
	f, r, err := pdf.Open(pdfPath)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}

	totalPages := r.NumPage()

	// If startPage is 0, start from page 1
	if startPage <= 0 {
		startPage = 1
	}

	// If endPage is 0 or greater than total pages, set to total pages
	if endPage <= 0 || endPage > totalPages {
		endPage = totalPages
	}

	for i := startPage; i <= endPage; i++ {
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}

		b, err := p.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		buf.WriteString(b)
	}
	return buf.String(), nil
}

// ExtractTextFromReader extracts text from a PDF provided as an io.Reader
func ExtractTextFromReader(reader io.Reader, startPage int, endPage int) (string, error) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "pdf-*.pdf")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Copy the reader content to the temp file
	if _, err := io.Copy(tempFile, reader); err != nil {
		return "", err
	}

	// Close the file to ensure all data is written
	tempFile.Close()

	// Extract text from the temp file
	return ExtractText(tempFile.Name(), startPage, endPage)
}
