package database

import (
	"database/sql"
	"fmt"
	"io"
	model "joshuamURD/wardens-court-summariser/internal/models"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type SQLITEDB struct {
	db             *sql.DB
	pdfStoragePath string
}

func initDB(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS decisions (citation TEXT PRIMARY KEY, date_delivered TEXT, court TEXT, warden TEXT, tenements TEXT, parties TEXT, summary TEXT)")
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pdf_documents (
		id TEXT PRIMARY KEY,
		citation TEXT,
		filename TEXT,
		filepath TEXT,
		upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (citation) REFERENCES decisions(citation)
	)`)
	if err != nil {
		return err
	}

	return nil
}

func NewSQLITEDB(dbPath string, pdfStoragePath string) (*SQLITEDB, error) {
	// Create PDF storage directory if it doesn't exist
	if err := os.MkdirAll(pdfStoragePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create PDF storage directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := initDB(db); err != nil {
		return nil, fmt.Errorf("failed to initialise database: %w", err)
	}

	sqliteDB := &SQLITEDB{
		db:             db,
		pdfStoragePath: pdfStoragePath,
	}

	return sqliteDB, nil
}

func (db *SQLITEDB) UploadFile(citation string, filename string, file io.Reader) (string, error) {
	id := uuid.New().String()

	// Create year-based directory structure
	currentTime := time.Now()
	yearDir := fmt.Sprintf("%d", currentTime.Year())
	dirPath := filepath.Join(db.pdfStoragePath, yearDir)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Create a unique filename to avoid collisions
	storedFilename := fmt.Sprintf("%s_%s", id, filename)
	filePath := filepath.Join(dirPath, storedFilename)

	// Create the file
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	// Copy the file content
	if _, err := io.Copy(outFile, file); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	// Store the file metadata in the database
	tx, err := db.db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Store relative path to make it more portable
	relativePath := filepath.Join(yearDir, storedFilename)

	_, err = tx.Exec(
		"INSERT INTO pdf_documents (id, citation, filename, filepath) VALUES (?, ?, ?, ?)",
		id, citation, filename, relativePath,
	)
	if err != nil {
		return "", fmt.Errorf("failed to insert PDF document metadata: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return id, nil
}

func (db *SQLITEDB) GetPDFDocument(citation string) ([]byte, string, error) {
	var filePath string
	var filename string

	err := db.db.QueryRow(
		"SELECT filepath, filename FROM pdf_documents WHERE citation = ?",
		citation,
	).Scan(&filePath, &filename)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", fmt.Errorf("no PDF document found for citation: %s", citation)
		}
		return nil, "", fmt.Errorf("failed to retrieve PDF document metadata: %w", err)
	}

	// Construct the full path
	fullPath := filepath.Join(db.pdfStoragePath, filePath)

	// Read the file content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read PDF file: %w", err)
	}

	return content, filename, nil
}

func (db *SQLITEDB) GetPDFPath(citation string) (string, string, error) {
	var filePath string
	var filename string

	err := db.db.QueryRow(
		"SELECT filepath, filename FROM pdf_documents WHERE citation = ?",
		citation,
	).Scan(&filePath, &filename)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("no PDF document found for citation: %s", citation)
		}
		return "", "", fmt.Errorf("failed to retrieve PDF document metadata: %w", err)
	}

	// Construct the full path
	fullPath := filepath.Join(db.pdfStoragePath, filePath)

	return fullPath, filename, nil
}

func (db *SQLITEDB) GetDecision(citation string) (*model.Decision, error) {
	return nil, nil
}

func (db *SQLITEDB) AllDecisions() ([]*model.Decision, error) {
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

func (db *SQLITEDB) GetDecisions() ([]model.Decision, error) {
	rows, err := db.db.Query("SELECT citation, date_delivered, court, warden, tenements, parties, summary FROM decisions")
	if err != nil {
		return nil, fmt.Errorf("failed to query decisions: %w", err)
	}
	defer rows.Close()

	var decisions []model.Decision
	for rows.Next() {
		var d model.Decision
		err := rows.Scan(&d.Citation, &d.DateDelivered, &d.Court, &d.Warden, &d.Tenements, &d.Parties, &d.Summary)
		if err != nil {
			return nil, fmt.Errorf("failed to scan decision: %w", err)
		}
		decisions = append(decisions, d)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating decisions: %w", err)
	}

	return decisions, nil
}
