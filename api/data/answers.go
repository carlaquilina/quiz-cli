package data

// Define a struct for a user's answer to a question
type UserAnswer struct {
	QuestionID int `json:"questionId"`
	AnswerID   int `json:"answerId"`
}
