package cmd

import (
	"fmt"
	"testing"
)

func TestValidateAnswer(t *testing.T) {
	// Get the arguments
	no := int64(4)
	answer := "manila"

	// Initialize SQLite3 database
	db, err := initDB()
	if err != nil {
		t.Errorf(fmt.Sprintf("error creating database and tables: %q", err.Error()))
		return
	}
	defer db.Close()

	// Create the question
	err = validateAnswer(db, no, false, answer)
	if err != nil {
		t.Errorf("error validating answer: %v", err.Error())
		return
	}
}
