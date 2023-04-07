package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleQuestions(t *testing.T) {
	req, err := http.NewRequest("GET", "/questions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleQuestions)

	handler.ServeHTTP(rr, req)

	// Check the response status code is 200
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response content type is JSON
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body contains the expected questions JSON
	expectedJSON := `[
		{
			"id": 1,
			"text": "What is the capital of France?",
			"answers": [
				{"id": 1, "text": "London", "correct": false},
				{"id": 2, "text": "Paris", "correct": true},
				{"id": 3, "text": "Berlin", "correct": false}
			]
		},
		{
			"id": 2,
			"text": "What is the largest continent?",
			"answers": [
				{"id": 1, "text": "Europe", "correct": false},
				{"id": 2, "text": "Asia", "correct": true},
				{"id": 3, "text": "Africa", "correct": false}
			]
		},
		{
			"id": 3,
			"text": "What is the smallest country in the world?",
			"answers": [
				{"id": 1, "text": "Vatican City", "correct": true},
				{"id": 2, "text": "Monaco", "correct": false},
				{"id": 3, "text": "San Marino", "correct": false}
			]
		}
	]`

	assert.JSONEq(t, expectedJSON, rr.Body.String())
}
