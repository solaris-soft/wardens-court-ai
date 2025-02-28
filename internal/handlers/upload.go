package handlers

import (
	"fmt"
	"io"
	model "joshuamURD/wardens-court-summariser/internal/models"
	views "joshuamURD/wardens-court-summariser/views/home"
	"joshuamURD/wardens-court-summariser/views/partials"
	"net/http"
)

type dataStore interface {
	uploadFile(bytes []byte) error
	getDecisions() ([]model.Decision, error)
}

func UploadFile(ds dataStore) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		fmt.Println("Uploading file")
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("decision")
		if err != nil {
			return Render(w, r, views.UploadStatus("error"))
		}
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			return Render(w, r, views.UploadStatus("error"))
		}
		// Return initial processing status
		return Render(w, r, partials.Table([]model.Decision{}))
	}
}
