package models

// Question is a struct that represents a question
type Question struct {
	No       int64  `json:"no"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
