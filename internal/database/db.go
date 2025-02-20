package database

import (
	"database/sql"
	model "joshuamURD/wardens-court-summariser/models"
	"os"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type SQLITEDB struct {
	db *sql.DB
}

func NewSQLITEDB(path string) (*SQLITEDB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return &SQLITEDB{db: db}, nil
}

func (db *SQLITEDB) UploadFile(file *os.File) error {
	return nil
}

func (db *SQLITEDB) GetDecision(citation string) (*model.Decision, error) {
	return nil, nil
}

func (db *SQLITEDB) GetAllDecisions() ([]*model.Decision, error) {
	return nil, nil
}

func (db *SQLITEDB) CreateDecision(decision *model.Decision) error {
	return nil
}

func (db *SQLITEDB) UpdateDecision(decision *model.Decision) error {
	return nil
}

func (db *SQLITEDB) DeleteDecision(id uuid.UUID) error {
	return nil
}
