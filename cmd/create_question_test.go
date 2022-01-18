package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestCreateQuestion(t *testing.T) {
	// Get the arguments
	no := int64(4)
	question := "What is the capital of Philippines?"
	answer := "Manila"

	// Initialize SQLite3 database
	db, err := initDB()
	if err != nil {
		t.Errorf(fmt.Sprintf("error creating database and tables: %q", err.Error()))
		return
	}
	defer db.Close()

	// Create the question
	err = createQuestion(db, no, question, answer)
	if err != nil {
		t.Errorf("error creating question: %v", err.Error())
		return
	}
	t.Log("Question created successfully!")
}

// initDB initializes the table questions
func initDB() (*sql.DB, error) {
	// Create database file
	os.Remove("../db/questions.db")

	db, err := sql.Open("sqlite3", "../db/questions.db")
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
