package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Define a struct for a question
type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

// Define a struct for an answer
type Answer struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

// Define a struct for a user's answer to a question
type UserAnswer struct {
	QuestionID int `json:"questionId"`
	AnswerID   int `json:"answerId"`
}

// Define a struct for the quiz statistics
type QuizStats struct {
	TotalQuizzesTaken int     `json:"totalQuizzesTaken"`
	UserRank          int     `json:"userRank"`
	UserScore         int     `json:"userScore"`
	AverageScore      float64 `json:"averageScore"`
	UserPercentage    float64 `json:"userPercentage"`
	Better            float64 `json:"better"`
}

var (
	currentUserName = "user1"
	// Define an array of questions
	questions = []Question{
		Question{
			ID:   1,
			Text: "What is the capital of France?",
			Answers: []Answer{
				Answer{ID: 1, Text: "London", Correct: false},
				Answer{ID: 2, Text: "Paris", Correct: true},
				Answer{ID: 3, Text: "Berlin", Correct: false},
			},
		},
		Question{
			ID:   2,
			Text: "What is the largest continent?",
			Answers: []Answer{
				Answer{ID: 1, Text: "Europe", Correct: false},
				Answer{ID: 2, Text: "Asia", Correct: true},
				Answer{ID: 3, Text: "Africa", Correct: false},
			},
		},
		Question{
			ID:   3,
			Text: "What is the smallest country in the world?",
			Answers: []Answer{
				Answer{ID: 1, Text: "Vatican City", Correct: true},
				Answer{ID: 2, Text: "Monaco", Correct: false},
				Answer{ID: 3, Text: "San Marino", Correct: false},
			},
		},
	}

	// Define a map to store the quiz results
	resultsMutex      sync.Mutex
	results           = map[string]int{"user2": 3, "user3": 1, "user4": 0} //map[username]score
	questionsAnswered = map[string][]int{}                                 //map[username]questionsAnswered
)

func isQuestionAnswered(username string, questionID int, questionsAnswered map[string][]int) bool {
	if answeredQuestions, ok := questionsAnswered[username]; ok {
		for _, q := range answeredQuestions {
			if q == questionID {
				return true
			}
		}
	}
	return false
}

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define the endpoints
	router.HandleFunc("/questions", handleQuestions).Methods("GET")
	router.HandleFunc("/answers", handleAnswer).Methods("POST")
	router.HandleFunc("/stats", handleStats).Methods("GET")

	// Serve the API
	log.Fatal(http.ListenAndServe(":8080", router))
}

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
func handleQuestions(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Marshal the questions to JSON and write it to the response body
	json.NewEncoder(w).Encode(questions)
}

// handleAnswer handles the POST /answers endpoint, which receives a list of answers for a quiz.
// The request body is a JSON-encoded array of objects, where each object represents an answer to a question.
// Each answer object has the following fields:
//   - questionID: the ID of the question that this answer is for (integer)
//   - answerID: the ID of the answer that the user selected (integer)
//
// The function returns a HTTP status code of 200 on success, and a HTTP status code of 400 on failure.
// Handler function for /answers
func handleAnswer(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to a UserAnswer struct
	var userAnswer UserAnswer
	err := json.NewDecoder(r.Body).Decode(&userAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the question and the selected answer
	var question *Question
	var selectedAnswer *Answer
	for i := range questions {
		if questions[i].ID == userAnswer.QuestionID {
			question = &questions[i]
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
		http.Error(w, "Invalid question or answer", http.StatusBadRequest)
		return
	}

	// If the question has already been answered, return an error
	if isQuestionAnswered(currentUserName, question.ID, questionsAnswered) {
		http.Error(w, "Question already answered", http.StatusBadRequest)
		return
	}

	// Add the question to the list of questions answered by the user
	questionsAnswered[currentUserName] = append(questionsAnswered[currentUserName], question.ID)

	if selectedAnswer.Correct {
		resultsMutex.Lock()
		results[currentUserName]++
		resultsMutex.Unlock()
	}

	// Return whether the selected answer is correct or not
	json.NewEncoder(w).Encode(selectedAnswer.Correct)
}

// handleStats handles the GET /stats endpoint, which returns statistics about the quiz results.
// The response body is a JSON-encoded object with the following fields:
//   - totalAttempts: the total number of times the quiz has been attempted (integer)
//   - averageScore: the average score of all users who have taken the quiz (float64)
//
// The function returns a HTTP status code of 200 on success, and a HTTP status code of 500 on failure.
// Handler function for /stats
func handleStats(w http.ResponseWriter, r *http.Request) {
	// Calculate the quiz statistics
	var totalQuizzesTaken int
	var userScore int
	var userRank int
	var totalScore int
	for username, result := range results {
		totalQuizzesTaken++
		totalScore += result
		if username == currentUserName {
			userScore = result
		}
	}
	for _, result := range results {
		if userScore > result {
			userRank++
		}
	}
	averageScore := float64(totalScore) / float64(totalQuizzesTaken)
	userPercentage := float64(userScore) / float64(len(questions)) * 100
	w.Header().Set("Content-Type", "application/json")
	// Return the quiz statistics
	json.NewEncoder(w).Encode(QuizStats{
		TotalQuizzesTaken: totalQuizzesTaken,
		UserRank:          userRank + 1,
		UserScore:         userScore,
		AverageScore:      averageScore,
		UserPercentage:    userPercentage,
		Better:            getBetterThan(results, currentUserName),
	})
}

func getBetterThan(results map[string]int, username string) float64 {
	numUsers := len(results)
	if numUsers == 0 {
		return 0
	}

	userScore, ok := results[username]
	if !ok {
		return 0
	}

	rank := 1
	for _, score := range results {
		if score > userScore {
			rank++
		}
	}

	percentBetter := float64(rank-1) / float64(numUsers) * 100
	return percentBetter
}
