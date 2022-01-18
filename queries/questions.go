package queries

import (
	"database/sql"
	"fmt"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/quiz_master/models"
)

// PutQuestion upserts question into the table questions
func PutQuestion(db *sql.DB, p models.Question) error {
	var stmt *sql.Stmt
	tx, err := db.Begin()
	if err != nil {
		itrlog.Fatalf("error beginning transaction: %q", err)
	}

	// Check if the question no already exists
	if IsQuestionNoExists(db, p.No) {
		// Update the question
		stmt, err = tx.Prepare("UPDATE questions SET question = ?, answer = ? WHERE no = ?")
		if err != nil {
			itrlog.Fatalf("error preparing statement: %q", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(p.Question, p.Answer, p.No)
		if err != nil {
			itrlog.Fatalf("error updating data: %q", err)
		}
	} else {
		// Insert the new question
		stmt, err = tx.Prepare("INSERT INTO questions(no, question, answer) values(?, ?, ?)")
		if err != nil {
			itrlog.Fatalf("error preparing statement: %q", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(p.No, p.Question, p.Answer)
		if err != nil {
			itrlog.Fatalf("error inserting data: %q", err)
		}
	}

	tx.Commit()
	return nil
}

// IsQuestionNoExists returns true if the question with the given no exists in the table questions
func IsQuestionNoExists(dbCon *sql.DB, no int64) bool {
	var id int64 = 0
	err := dbCon.QueryRow(`SELECT no FROM questions WHERE no = ?`, no).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false // returned no rows
		}
		return false
	}
	return true
}

// DeleteQuestion deletes the question with the given no from the table questions
func DeleteQuestion(db *sql.DB, no int64) error {
	tx, err := db.Begin()
	if err != nil {
		itrlog.Fatalf("error beginning transaction: %q", err)
	}
	stmt, err := tx.Prepare("DELETE FROM questions WHERE no = ?")
	if err != nil {
		itrlog.Fatalf("error preparing statement: %q", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(no)
	if err != nil {
		itrlog.Fatalf("error deleting data: %q", err)
	}

	tx.Commit()
	return nil
}

// GetQuestions returns all questions from the table questions
func GetQuestions(db *sql.DB) ([]models.Question, error) {
	qResults, err := db.Query("SELECT no, question, answer FROM questions ORDER BY no")
	if err != nil {
		itrlog.Fatalf("error querying data: %q", err)
	}
	defer qResults.Close()

	var results []models.Question
	for qResults.Next() {
		var q models.Question
		err = qResults.Scan(&q.No, &q.Question, &q.Answer)
		if err != nil {
			errInfo := fmt.Sprintf("error scanning questions: %s", err.Error())
			itrlog.Errorf(errInfo)
			return nil, err
		}
		results = append(results, q)
	}
	return results, nil
}

// GetQuestion returns the question with the given no from the table questions
func GetQuestion(db *sql.DB, no int64) ([]models.Question, error) {
	qResults, err := db.Query("SELECT no, question, answer FROM questions WHERE no = ?", no)
	if err != nil {
		itrlog.Fatalf("error querying data: %q", err)
	}
	defer qResults.Close()

	var results []models.Question
	for qResults.Next() {
		var q models.Question
		err = qResults.Scan(&q.No, &q.Question, &q.Answer)
		if err != nil {
			errInfo := fmt.Sprintf("error scanning question: %s", err.Error())
			itrlog.Errorf(errInfo)
			return nil, err
		}
		results = append(results, q)
	}
	return results, nil
}
