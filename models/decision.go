package model

import (
	"time"
)

// Party represents a party in the Wardens Court
type Party struct {
	Name string
	Role string // Applicant, Respondent, Objector
}

// Decision represents a decision made in the Wardens Court
type Decision struct {
	DecisionNumber string    `json:"decision_number"`
	DateDelivered  time.Time `json:"date_delivered"`
	Court          string    `json:"court"`
	Warden         string    `json:"warden"`
	Tenements      []string  `json:"tenements"`
	Parties        string    `json:"parties"`
	Summary        string    `json:"summary"`
	DocumentURL    string    `json:"document_url"`
}
