package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spf13/myapp/api/data"
	"github.com/spf13/myapp/api/models"
)

func TestHandleAnswer(t *testing.T) {
	data.Questions = []models.Question{
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
	// Create a test HTTP server
	testServer := httptest.NewServer(http.HandlerFunc(HandleAnswer))
	defer testServer.Close()

	// Test cases
	cases := []struct {
		description        string
		request            models.UserAnswer
		expectedResponse   string
		expectedStatusCode int
		incorrectBody      bool
	}{
		{
			description: "Correct answer",
			request: models.UserAnswer{
				QuestionID: 1,
				AnswerID:   2,
			},
			expectedResponse:   "true\n",
			expectedStatusCode: http.StatusOK,
		},
		{
			description: "Incorrect answer",
			request: models.UserAnswer{
				QuestionID: 1,
				AnswerID:   1,
			},
			expectedResponse:   "false\n",
			expectedStatusCode: http.StatusOK,
		},
		{
			description: "Invalid question",
			request: models.UserAnswer{
				QuestionID: 4,
				AnswerID:   2,
			},
			expectedResponse:   "{\"error\":\"Invalid question or answer\"}\n",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description: "Invalid answer",
			request: models.UserAnswer{
				QuestionID: 1,
				AnswerID:   4,
			},
			expectedResponse:   "{\"error\":\"Invalid question or answer\"}\n",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "Invalid body",
			expectedResponse:   "{\"Could not decode body - error\":\"EOF\"}\n",
			expectedStatusCode: http.StatusBadRequest,
			incorrectBody:      true,
		},
	}

	// Perform tests
	for _, tc := range cases {
		data.QuestionsAnswered = models.QuestionAnswer{}
		data.Results = models.Results{
			M: map[string]int{"user2": 3, "user3": 1, "user4": 0, "user5": 0},
		}
		// Encode the request body to JSON
		requestBody, err := json.Marshal(tc.request)
		assert.NoError(t, err)

		// Send the request to the test server
		res := &http.Response{}
		err = nil
		if tc.incorrectBody {
			res, err = http.Post(testServer.URL, "application/json", bytes.NewBuffer([]byte("")))
		} else {
			res, err = http.Post(testServer.URL, "application/json", bytes.NewBuffer(requestBody))
		}
		assert.NoError(t, err)
		defer res.Body.Close()

		// Check the response status code
		assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

		bytes, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		// Check the expected response
		assert.Equal(t, tc.expectedResponse, string(bytes))
	}
}
