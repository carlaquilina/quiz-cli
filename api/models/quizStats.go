package models

// Define a struct for the quiz statistics
type QuizStats struct {
	TotalQuizzesTaken int     `json:"totalQuizzesTaken"` //the total amount of quizez taken by all users
	UserRank          int     `json:"userRank"`          //the higher the value the better the rank
	UserScore         int     `json:"userScore"`         //the score of the current user requesting stats
	AverageScore      float64 `json:"averageScore"`      //average score of all users
	UserPercentage    float64 `json:"userPercentage"`    //the percentage of the current user's score out of the total amount of score(marks)
	Better            float64 `json:"better"`            //the percentage of users that the current user is better than
}
