package cmd

import (
	"fmt"
	"testing"
)

func TestGetQuestions(t *testing.T) {
	// Initialize SQLite3 database
	db, err := initDB()
	if err != nil {
		t.Errorf(fmt.Sprintf("error creating database and tables: %q", err.Error()))
		return
	}
	defer db.Close()

	// Get the questions
	err = getQuestions(db)
	if err != nil {
		t.Errorf("error getting questions: %v", err.Error())
		return
	}
	t.Log("Questions retrieved successfully!")
}
