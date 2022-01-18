package words2num

import (
	"errors"
	"strings"
)

const (
	ErrNotPureWord = "INVALID_NUM_WORD"
)

// ConvertWordsToNumber converts a string number to number
func ConvertWordsToNumber(words string) (int, error) {
	words = strings.ToLower(words)
	words = strings.TrimSpace(words)
	words = strings.Replace(words, " ", "", -1)
	words = strings.Replace(words, ",", "", -1)
	words = strings.Replace(words, "and", "", -1)

	var number int
	for _, word := range strings.Split(words, " ") {
		if value, ok := WordsValue[word]; ok {
			number = number*10 + value
		} else {
			return 0, errors.New(ErrNotPureWord)
		}
	}
	return number, nil
}
