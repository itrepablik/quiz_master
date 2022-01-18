/*
Copyright © 2021 ITRepablik itrepablik@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/quiz_master/models"
	"github.com/itrepablik/quiz_master/queries"
	"github.com/spf13/cobra"
)

// updateQuestionCmd represents the update_question command
var updateQuestionCmd = &cobra.Command{
	Use:   "update_question",
	Short: "Updates a question",
	Long:  `Updates a question with the following required parameters: <no> <question> <answer>`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the arguments
		no, err := strconv.Atoi(args[0])
		if err != nil {
			errMsg := fmt.Sprintf("error converting no to integer: %v", err.Error())
			itrlog.Error(errMsg)
			fmt.Println(errMsg)
			return
		}

		question := args[1]
		answer := args[2]

		// Prepare the payload
		payload := models.Question{
			No:       int64(no),
			Question: question,
			Answer:   answer,
		}

		// Initialize SQLite3 database
		db, err := queries.InitTableQuestions()
		if err != nil {
			errMsg := fmt.Sprintf("error creating database and tables: %q", err.Error())
			itrlog.Fatalf(errMsg)
		}
		defer db.Close()

		// Update the question
		err = updateQuestion(db, payload)
		if err != nil {
			errMsg := fmt.Sprintf("error updating question: %v", err.Error())
			itrlog.Error(errMsg)
			fmt.Println(errMsg)
			return
		}

		fmt.Println("Question updated successfully!")
	},
}

// updateQuestion updates a question
func updateQuestion(db *sql.DB, p models.Question) error {
	// Check if the question no already exists
	if !queries.IsQuestionNoExists(db, p.No) {
		return fmt.Errorf("question no %d is not exists", p.No)
	}

	// Update the new question
	err := queries.PutQuestion(db, p)
	if err != nil {
		return fmt.Errorf("error updating question: %q", err.Error())
	}
	return nil
}

func init() {
	rootCmd.AddCommand(updateQuestionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateQuestionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateQuestionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
