package cmd

import "fmt"

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

func (q Question) Print() {
	fmt.Printf("Question %d: %s\n", q.ID, q.Text)
	for _, answer := range q.Answers {
		fmt.Printf("%d. %s\n", answer.ID, answer.Text)
	}
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
