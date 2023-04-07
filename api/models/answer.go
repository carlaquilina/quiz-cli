package models

// Define a struct for an answer
type Answer struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}
