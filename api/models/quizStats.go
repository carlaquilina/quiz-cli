package models

// Define a struct for the quiz statistics
type QuizStats struct {
	TotalQuizzesTaken int     `json:"totalQuizzesTaken"`
	UserRank          int     `json:"userRank"`
	UserScore         int     `json:"userScore"`
	AverageScore      float64 `json:"averageScore"`
	UserPercentage    float64 `json:"userPercentage"`
	Better            float64 `json:"better"`
}
