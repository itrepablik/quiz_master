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
	"strings"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/quiz_master/queries"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// deleteQuestionCmd represents the delete_question command
var deleteQuestionCmd = &cobra.Command{
	Use:   "delete_question",
	Short: "Delete a question",
	Long:  `Delete a question with the following required parameter: <no>`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the arguments
		no, err := strconv.Atoi(args[0])
		if err != nil {
			errMsg := fmt.Sprintf("error converting no to integer: %v", err.Error())
			itrlog.Error(errMsg)
			fmt.Println(errMsg)
			return
		}

		// Initialize SQLite3 database
		db, err := queries.InitTableQuestions()
		if err != nil {
			errMsg := fmt.Sprintf("error initializing database: %v", err.Error())
			itrlog.Error(errMsg)
			fmt.Println(errMsg)
			return
		}
		defer db.Close()

		// Prompt user to confirm
		prompt := promptui.Prompt{
			Label:     "Delete Question No " + strconv.Itoa(no),
			IsConfirm: true,
		}

		confirmOpt, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose %q\n", confirmOpt)

		if strings.ToLower(strings.TrimSpace(confirmOpt)) == "y" || strings.ToLower(strings.TrimSpace(confirmOpt)) == "yes" {
			// Delete the question by no
			err := delQuestion(db, int64(no))
			if err != nil {
				errMsg := fmt.Sprintf("error deleting question: %q", err.Error())
				itrlog.Errorf(errMsg)
				fmt.Println(errMsg)
				return
			}
			delInfo := fmt.Sprintf("Question no %d was deleted", no)
			fmt.Println(delInfo)
			itrlog.Warnf(delInfo)
		}
	},
}

// delQuestion is a function to delete a question
func delQuestion(db *sql.DB, no int64) error {
	// Check if the question no already exists
	if !queries.IsQuestionNoExists(db, int64(no)) {
		return fmt.Errorf("question no %d is not exists", no)
	}

	// Delete the question by no
	err := queries.DeleteQuestion(db, int64(no))
	if err != nil {
		return fmt.Errorf("error deleting question: %q", err.Error())
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deleteQuestionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteQuestionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteQuestionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
