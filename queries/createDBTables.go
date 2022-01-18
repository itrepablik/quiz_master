package queries

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

// InitTableQuestions initializes the table questions
func InitTableQuestions() (*sql.DB, error) {
	// Create database file
	os.Remove("./db/questions.db")

	db, err := sql.Open("sqlite3", "./db/questions.db")
	if err != nil {
		return db, fmt.Errorf("error opening database: %q", err)
	}

	// Create table if not exists questions
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS questions (no integer not null primary key, question text, answer text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return db, fmt.Errorf("%q: %s", err, sqlStmt)
	}
	return db, nil
}
