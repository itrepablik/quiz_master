package cmd

import (
	"fmt"
	"testing"
)

func TestDelQuestion(t *testing.T) {
	// Get the arguments
	no := int64(4)

	// Initialize SQLite3 database
	db, err := initDB()
	if err != nil {
		t.Errorf(fmt.Sprintf("error creating database and tables: %q", err.Error()))
		return
	}
	defer db.Close()

	// Delete the question
	err = delQuestion(db, no)
	if err != nil {
		t.Errorf("error deleting question: %v", err.Error())
		return
	}
	t.Log("Question deleted successfully!")
}
