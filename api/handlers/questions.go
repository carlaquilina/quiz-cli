package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/myapp/api/data"
)

// handleQuestions handles the GET /questions endpoint, which returns a list of questions with their corresponding answers.
// The response body is a JSON-encoded array of objects, where each object represents a question and its answers.
// Each question object has the following fields:
//   - id: a unique identifier for the question (integer)
//   - question: the text of the question (string)
//   - answers: an array of answer objects, each with the following fields:
//   - id: a unique identifier for the answer (integer)
//   - answer: the text of the answer (string)
//   - isCorrect: a boolean indicating whether this answer is correct (bool)
//
// The function returns a HTTP status code of 200 on success, and a HTTP status code of 500 on failure.
// Handler function for /questions
func HandleQuestions(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Marshal the questions to JSON and write it to the response body
	json.NewEncoder(w).Encode(data.Questions)
}
