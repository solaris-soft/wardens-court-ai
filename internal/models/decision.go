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
	Citation      string    `json:"citation"`
	DateDelivered time.Time `json:"date_delivered"`
	Court         string    `json:"court"`
	Warden        string    `json:"warden"`
	Tenements     []string  `json:"tenements"`
	Parties       [2]Party  `json:"parties"`
	Summary       string    `json:"summary"`
	DocumentURL   string    `json:"document_url"`
}

type DecisionStore struct {
	Decisions map[string]Decision
}

func (ds *DecisionStore) AddDecision(decision Decision) {
	ds.Decisions[decision.Citation] = decision
}
