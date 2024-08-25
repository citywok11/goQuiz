package api

import (
	"bytes"
	"encoding/json"
	"goQuiz/internal/models"
	"net/http"
)

var BaseURL = "http://localhost:8081"

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL:    BaseURL,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) GetQuestions() ([]models.Question, error) {
	resp, err := c.HTTPClient.Get(c.BaseURL + "/questions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var questions []models.Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	return questions, err
}

func (c *Client) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	data, err := json.Marshal(answers)
	if err != nil {
		return models.QuizResult{}, err
	}

	resp, err := c.HTTPClient.Post(c.BaseURL+"/submit", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return models.QuizResult{}, err
	}
	defer resp.Body.Close()

	var result models.QuizResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// Compatibility functions for existing code
func GetQuestions() ([]models.Question, error) {
	client := NewClient()
	return client.GetQuestions()
}

func SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	client := NewClient()
	return client.SubmitAnswers(answers)
}
