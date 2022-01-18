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

// createQuestionCmd represents the create_question command
var createQuestionCmd = &cobra.Command{
	Use:   "create_question",
	Short: "Creates a question",
	Long:  `Creates a question with the following required parameters: <no> <question> <answer>`,
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

		// Initialize SQLite3 database
		db, err := queries.InitTableQuestions()
		if err != nil {
			errMsg := fmt.Sprintf("error initializing database: %v", err.Error())
			itrlog.Error(errMsg)
			fmt.Println(errMsg)
			return
		}
		defer db.Close()

		// Create the question
		err = createQuestion(db, int64(no), question, answer)
		if err != nil {
			errMsg := fmt.Sprintf("error creating question: %v", err.Error())
			itrlog.Error(errMsg)
			fmt.Println(errMsg)
			return
		}
		fmt.Println("Question created successfully!")
	},
}

// createQuestion creates a question
func createQuestion(db *sql.DB, no int64, question, answer string) error {
	// Prepare the payload
	payload := models.Question{
		No:       no,
		Question: question,
		Answer:   answer,
	}

	// Check if the question no already exists
	if queries.IsQuestionNoExists(db, int64(no)) {
		return fmt.Errorf(fmt.Sprintf("Question no %d already exists", no))
	}

	// Create the new question
	err := queries.PutQuestion(db, payload)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("error creating question: %q", err.Error()))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(createQuestionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createQuestionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createQuestionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
