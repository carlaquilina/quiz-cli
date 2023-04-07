package models_test

import (
	"testing"

	"github.com/spf13/myapp/api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetBetterThan(t *testing.T) {
	cases := []struct {
		description     string
		results         *models.Results
		expectedResults float64
		username        string
	}{
		{
			description: "Empty Results",
			username:    "Alice",
			results: &models.Results{
				M: map[string]int{},
			},
			expectedResults: 0,
		},
		{
			description: "Test Results",
			username:    "Alice",
			results: &models.Results{
				M: map[string]int{
					"Alice": 1,
					"Bob":   2,
				},
			},
			expectedResults: 0,
		},
		{
			description: "Equal Test Results",
			username:    "Alice",
			results: &models.Results{
				M: map[string]int{
					"Alice": 1,
					"Bob":   1,
				},
			},
			expectedResults: 0,
		},
		{
			description: "Results wiuth 4 users",
			username:    "Alice",
			results: &models.Results{
				M: map[string]int{
					"Alice":  1,
					"Bob":    0,
					"George": 2,
					"Kay":    2,
				},
			},
			expectedResults: 33.33333333333333,
		},
		{
			description: "Username not in results",
			username:    "not in result",
			results: &models.Results{
				M: map[string]int{
					"Alice":  1,
					"Bob":    0,
					"George": 2,
					"Kay":    2,
				},
			},
			expectedResults: 0,
		},
	}
	for _, tc := range cases {
		result := tc.results.GetBetterThan(tc.username)
		assert.Equal(t, tc.expectedResults, result, tc.description)

	}
}

func TestIncreaseScore(t *testing.T) {
	cases := []struct {
		description string
		username    string
		results     *models.Results
		expected    map[string]int
	}{
		{
			description: "Increase score for existing user",
			username:    "Alice",
			results: &models.Results{
				M: map[string]int{"Alice": 1},
			},
			expected: map[string]int{"Alice": 2},
		},
		{
			description: "Increase score for new user",
			username:    "Bob",
			results: &models.Results{
				M: map[string]int{"Alice": 1},
			},
			expected: map[string]int{"Alice": 1, "Bob": 1},
		},
	}

	for _, tc := range cases {
		tc.results.IncreaseScore(tc.username)
		assert.Equal(t, tc.expected, tc.results.M)
	}
}
