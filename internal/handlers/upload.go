package handlers

import (
	"fmt"
	model "joshuamURD/wardens-court-summariser/models"
	"joshuamURD/wardens-court-summariser/views/partials"
	"net/http"
	"os"
	"time"
)

type dataStore interface {
	UploadFile(file *os.File) error
}

func UploadFile(ds dataStore) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		fmt.Println("Uploading file")
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("decision")
		if err != nil {
			return fmt.Errorf("failed to get file: %w", err)
		}
		defer file.Close()

		decision := model.Decision{
			Citation:      "[2024] WAMW 1",
			DateDelivered: time.Now(),
			Court:         "Wardens Court",
			Warden:        "Warden McPhee",
			Tenements:     []string{"E45/1234", "E45/1235"},
			Parties:       [2]model.Party{{Name: "John Doe", Role: "Applicant"}, {Name: "Jane Doe", Role: "Respondent"}},
			Summary:       "This is a summary of the decision",
			DocumentURL:   "https://example.com/decision.pdf",
		}

		return Render(w, r, partials.TableRow(decision))
	}
}
