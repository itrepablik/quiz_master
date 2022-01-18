package cmd

import (
	"fmt"
	"testing"

	"github.com/itrepablik/quiz_master/models"
)

func TestUpdateQuestion(t *testing.T) {
	// Get the arguments
	no := int64(4)
	question := "What is the capital of Philippines?"
	answer := "Manila"

	p := models.Question{
		No:       no,
		Question: question,
		Answer:   answer,
	}

	// Initialize SQLite3 database
	db, err := initDB()
	if err != nil {
		t.Errorf(fmt.Sprintf("error creating database and tables: %q", err.Error()))
		return
	}
	defer db.Close()

	// Update the question
	err = updateQuestion(db, p)
	if err != nil {
		t.Errorf("error updating question: %v", err.Error())
		return
	}
	t.Log("Question updated successfully!")
}
