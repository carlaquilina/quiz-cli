package data

import "github.com/spf13/myapp/api/models"

var (
	Questions = []models.Question{
		models.Question{
			ID:   1,
			Text: "What is the capital of France?",
			Answers: []models.Answer{
				models.Answer{ID: 1, Text: "London", Correct: false},
				models.Answer{ID: 2, Text: "Paris", Correct: true},
				models.Answer{ID: 3, Text: "Berlin", Correct: false},
			},
		},
		models.Question{
			ID:   2,
			Text: "What is the largest continent?",
			Answers: []models.Answer{
				models.Answer{ID: 1, Text: "Europe", Correct: false},
				models.Answer{ID: 2, Text: "Asia", Correct: true},
				models.Answer{ID: 3, Text: "Africa", Correct: false},
			},
		},
		models.Question{
			ID:   3,
			Text: "What is the smallest country in the world?",
			Answers: []models.Answer{
				models.Answer{ID: 1, Text: "Vatican City", Correct: true},
				models.Answer{ID: 2, Text: "Monaco", Correct: false},
				models.Answer{ID: 3, Text: "San Marino", Correct: false},
			},
		},
	}
)
