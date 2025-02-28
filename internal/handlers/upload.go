package handlers

import (
	"fmt"
	pdfReader "joshuamURD/wardens-court-summariser/internal/pdf"
	"net/http"
)

func UploadFile(ds dataStore) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		fmt.Println("Uploading file")

		// Parse the multipart form
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			return fmt.Errorf("failed to parse form: %w", err)
		}

		// Get the file from the form
		file, _, err := r.FormFile("decision")
		if err != nil {
			return fmt.Errorf("failed to get file: %w", err)
		}
		defer file.Close()

		// Extract text from the PDF
		text, err := pdfReader.ExtractTextFromReader(file, 1, 1)
		if err != nil {
			return fmt.Errorf("failed to extract text from PDF: %w", err)
		}

		// Reset the file for upload to database
		if _, err := file.Seek(0, 0); err != nil {
			return fmt.Errorf("failed to reset file: %w", err)
		}

		fmt.Println(text)

		// Return the table with decisions
		return nil
	}
}
