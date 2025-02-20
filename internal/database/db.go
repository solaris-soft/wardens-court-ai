package database

import (
	"database/sql"
	"fmt"
	model "joshuamURD/wardens-court-summariser/models"
	"os"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type SQLITEDB struct {
	db *sql.DB
}

func initDB(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS decisions (citation TEXT PRIMARY KEY, date_delivered TEXT, court TEXT, warden TEXT, tenements TEXT, parties TEXT, summary TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func NewSQLITEDB(path string) (*SQLITEDB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	var sqliteDB *SQLITEDB
	if err := initDB(db); err != nil {
		return nil, fmt.Errorf("failed to initialise database: %w", err)
	}
	sqliteDB = &SQLITEDB{db: db}
	return sqliteDB, nil
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
