package handlers

import (
	model "joshuamURD/wardens-court-summariser/internal/models"
	views "joshuamURD/wardens-court-summariser/views/home"
	"joshuamURD/wardens-court-summariser/views/partials"
	"net/http"
)

type dataStore interface {
	AllDecisions() ([]model.Decision, error)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Index([]model.Decision{}))
}

func HandleTable(ds dataStore) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		decisions, err := ds.AllDecisions()
		if err != nil {
			return err
		}
		return Render(w, r, partials.Table(decisions))
	}
}
