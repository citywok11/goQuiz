package client

import (
	"encoding/json"
	"goQuiz/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client)
	assert.NotEmpty(t, client.BaseURL)
}

func TestClient_GetQuestions(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/questions", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		questions := []models.Question{
			{ID: 1, Text: "Test Question 1"},
			{ID: 2, Text: "Test Question 2"},
		}
		json.NewEncoder(w).Encode(questions)
	}))
	defer server.Close()

	// Create a client with the test server URL
	client := &Client{BaseURL: server.URL}

	// Test the GetQuestions method
	questions, err := client.GetQuestions()
	assert.NoError(t, err)
	assert.Len(t, questions, 2)
	assert.Equal(t, "Test Question 1", questions[0].Text)
	assert.Equal(t, "Test Question 2", questions[1].Text)
}

func TestClient_GetQuestions_Error(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Create a client with the test server URL
	client := &Client{BaseURL: server.URL}

	questions, err := client.GetQuestions()
	assert.Error(t, err)
	assert.Nil(t, questions)
	assert.Contains(t, err.Error(), "EOF")
}

func TestClient_SubmitAnswers(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/submit", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var answers []models.UserAnswer
		json.NewDecoder(r.Body).Decode(&answers)
		assert.Len(t, answers, 2)

		result := models.QuizResult{
			CorrectAnswers: 1,
			TotalQuestions: 2,
			Percentile:     50.0,
		}
		json.NewEncoder(w).Encode(result)
	}))
	defer server.Close()

	// Create a client with the test server URL
	client := &Client{BaseURL: server.URL}

	// Test the SubmitAnswers method
	answers := []models.UserAnswer{
		{QuestionID: 1, AnswerIndex: 0},
		{QuestionID: 2, AnswerIndex: 1},
	}
	result, err := client.SubmitAnswers(answers)
	assert.NoError(t, err)
	assert.Equal(t, 1, result.CorrectAnswers)
	assert.Equal(t, 2, result.TotalQuestions)
	assert.Equal(t, 50.0, result.Percentile)
}

func TestClient_SubmitAnswers_Error(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Create a client with the test server URL
	client := &Client{BaseURL: server.URL}

	// Test the SubmitAnswers method with an error response
	answers := []models.UserAnswer{
		{QuestionID: 1, AnswerIndex: 0},
	}
	result, err := client.SubmitAnswers(answers)
	assert.Error(t, err)
	assert.Equal(t, models.QuizResult{}, result)
	assert.Contains(t, err.Error(), "EOF")
}
