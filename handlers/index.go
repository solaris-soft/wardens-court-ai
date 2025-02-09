package handlers

import (
	model "joshuamURD/wardens-court-summariser/models"
	views "joshuamURD/wardens-court-summariser/views/home"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Index([]model.Decision{}))
}
