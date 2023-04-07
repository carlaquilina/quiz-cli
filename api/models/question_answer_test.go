package models_test

import (
	"testing"

	"github.com/spf13/myapp/api/models"
	"github.com/stretchr/testify/assert"
)

func TestIsQuestionAnswered(t *testing.T) {
	cases := []struct {
		description string
		questionID  int
		username    string
		answers     models.QuestionAnswer
		expected    bool
	}{
		{
			description: "User has not answered any questions",
			questionID:  1,
			username:    "Alice",
			answers:     models.QuestionAnswer{},
			expected:    false,
		},
		{
			description: "User has answered the question",
			questionID:  2,
			username:    "Alice",
			answers: models.QuestionAnswer{
				"Alice": []int{1, 2, 3},
			},
			expected: true,
		},
		{
			description: "User has not answered the question",
			questionID:  2,
			username:    "Alice",
			answers: models.QuestionAnswer{
				"Alice": []int{1, 3},
			},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			result := tc.answers.IsQuestionAnswered(tc.username, tc.questionID)

			assert.Equal(t, tc.expected, result)
		})
	}
}
