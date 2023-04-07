package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/myapp/api/data"
	"github.com/spf13/myapp/api/models"
)

// handleAnswer handles the POST /answers endpoint, which receives a list of answers for a quiz.
// The request body is a JSON-encoded array of objects, where each object represents an answer to a question.
// Each answer object has the following fields:
//   - questionID: the ID of the question that this answer is for (integer)
//   - answerID: the ID of the answer that the user selected (integer)
//
// The function returns a HTTP status code of 200 on success, and a HTTP status code of 400 on failure.
// Handler function for /answers
func HandleAnswer(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to a UserAnswer struct
	var userAnswer models.UserAnswer
	err := json.NewDecoder(r.Body).Decode(&userAnswer)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Could not decode body - error": err.Error()})
		return
	}

	// Find the question and the selected answer
	var question *models.Question
	var selectedAnswer *models.Answer
	for _, q := range data.Questions {
		if q.ID == userAnswer.QuestionID {
			question = &q
			for j := range question.Answers {
				if question.Answers[j].ID == userAnswer.AnswerID {
					selectedAnswer = &question.Answers[j]
					break
				}
			}
			break
		}
	}

	// If the question or the answer is not found, return an error
	if question == nil || selectedAnswer == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid question or answer"})
		return
	}

	// If the question has already been answered, return an error
	if data.QuestionsAnswered.IsQuestionAnswered(data.CurrentUserName, question.ID) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Question already answered"})
		return
	}

	// Add the question to the list of questions answered by the user
	data.QuestionsAnswered[data.CurrentUserName] = append(data.QuestionsAnswered[data.CurrentUserName], question.ID)

	if selectedAnswer.Correct {
		data.Results.IncreaseScore(data.CurrentUserName)
	}

	// Return whether the selected answer is correct or not
	json.NewEncoder(w).Encode(selectedAnswer.Correct)
}
