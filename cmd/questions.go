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

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/quiz_master/queries"
	"github.com/spf13/cobra"
)

// QuestionsCmd represents the question command
var QuestionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "Shows question list",
	Long:  `Shows all the questions in a list`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize SQLite3 database
		db, err := queries.InitTableQuestions()
		if err != nil {
			errMsg := fmt.Sprintf("error creating database and tables: %q", err.Error())
			itrlog.Fatalf(errMsg)
		}
		defer db.Close()

		// Get all the questions sorted by no in ascending order
		err = getQuestions(db)
		if err != nil {
			errMsg := fmt.Sprintf("error retrieving questions: %q", err.Error())
			itrlog.Errorf(errMsg)
			fmt.Println(errMsg)
			return
		}
	},
}

// getQuestions returns all the questions in a list
func getQuestions(db *sql.DB) error {
	// Get all the questions sorted by no in ascending order
	results, err := queries.GetQuestions(db)
	if err != nil {
		return fmt.Errorf("error retrieving questions: %q", err.Error())
	}
	fmt.Println("No.\tQuestion\tAnswer")
	for _, result := range results {
		fmt.Printf("%d\t%s\t%s\n", result.No, result.Question, result.Answer)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(QuestionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// QuestionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// QuestionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
