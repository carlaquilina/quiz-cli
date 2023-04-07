package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/myapp/api/data"
	"github.com/spf13/myapp/api/models"
)

// handleStats handles the GET /stats endpoint, which returns statistics about the quiz results.
// The response body is a JSON-encoded object with the following fields:
//   - totalAttempts: the total number of times the quiz has been attempted (integer)
//   - averageScore: the average score of all users who have taken the quiz (float64)
//
// The function returns a HTTP status code of 200 on success, and a HTTP status code of 500 on failure.
// Handler function for /stats
func HandleStats(w http.ResponseWriter, r *http.Request) {
	// Calculate the quiz statistics
	var totalQuizzesTaken int
	var userScore int
	var userRank int
	var totalScore int
	for username, result := range data.Results.M {
		totalQuizzesTaken++
		totalScore += result
		if username == data.CurrentUserName {
			userScore = result
		}
	}
	for _, result := range data.Results.M {
		if userScore > result {
			userRank++
		}
	}
	averageScore := float64(totalScore) / float64(totalQuizzesTaken)
	userPercentage := float64(userScore) / float64(len(data.Questions)) * 100
	w.Header().Set("Content-Type", "application/json")
	// Return the quiz statistics
	json.NewEncoder(w).Encode(models.QuizStats{
		TotalQuizzesTaken: totalQuizzesTaken,
		UserRank:          userRank + 1,
		UserScore:         userScore,
		AverageScore:      averageScore,
		UserPercentage:    userPercentage,
		Better:            data.Results.GetBetterThan(data.CurrentUserName),
	})
}
