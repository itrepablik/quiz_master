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

	"github.com/fatih/color"
	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/quiz_master/queries"
	"github.com/itrepablik/quiz_master/utils/words2num"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// QuestionCmd represents the question command
var QuestionCmd = &cobra.Command{
	Use:   "question",
	Short: "Shows a question",
	Long:  `Shows a question with the following required parameter: <no>`,
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
			errMsg := fmt.Sprintf("error creating database and tables: %q", err.Error())
			itrlog.Fatalf(errMsg)
		}
		defer db.Close()

		err = validateAnswer(db, int64(no), true, "")
		if err != nil {
			errMsg := fmt.Sprintf("error validating answer: %q", err.Error())
			itrlog.Error(errMsg)
			fmt.Println(errMsg)
			return
		}
	},
}

// validateAnswer checks if the answer is a number or a word
func validateAnswer(db *sql.DB, no int64, isPrompt bool, resTest string) error {
	// Check if the question no already exists
	if !queries.IsQuestionNoExists(db, no) {
		return fmt.Errorf("question no %d is not exists", no)
	}

	// Get all the questions sorted by no in ascending order
	results, err := queries.GetQuestion(db, no)
	if err != nil {
		return fmt.Errorf("error retrieving questions: %q", err.Error())
	}

	loadQuestion := ""
	if len(results) > 0 {
		loadQuestion = results[0].Question
	}

	// Ask the user for the answer
	isAnsNumeric := true
	result := ""
	if isPrompt {
		validate := func(input string) error {
			if strings.TrimSpace(input) == "" {
				return fmt.Errorf("answer is required")
			}
			return nil
		}

		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		}

		prompt := promptui.Prompt{
			Label:     loadQuestion,
			Templates: templates,
			Validate:  validate,
		}

		result, err = prompt.Run()

		if err != nil {
			return fmt.Errorf("prompt failed %v", err)
		}

		// Check if the answer entered is a number
		_, err = strconv.Atoi(result)
		if err != nil {
			isAnsNumeric = false
		}
	}

	// For unit testing
	if resTest != "" {
		result = resTest
	}

	// Check if the answer is correct
	if isAnsNumeric {
		// The answer is a numerical value
		if strings.EqualFold(strings.TrimSpace(result), strings.TrimSpace(results[0].Answer)) {
			color.Green("CORRECT!\n")
		} else {
			color.Red("INCORRECT!\n")
		}
	} else {
		// Convert numbers to words
		isPureWord := true
		num, err := words2num.ConvertWordsToNumber(strings.TrimSpace(result))
		if err != nil {
			if err.Error() == words2num.ErrNotPureWord {
				isPureWord = false
			}
		}

		if isPureWord {
			// Convert num to string and compare with answer
			if strings.EqualFold(strings.TrimSpace(strconv.Itoa(num)), strings.TrimSpace(results[0].Answer)) {
				color.Green("CORRECT!\n")
			} else {
				color.Red("INCORRECT!\n")
			}
		} else {
			// Possible is a plain word answer
			if strings.EqualFold(strings.TrimSpace(result), strings.TrimSpace(results[0].Answer)) {
				color.Green("CORRECT!\n")
			} else {
				color.Red("INCORRECT!\n")
			}
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(QuestionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// QuestionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// QuestionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
