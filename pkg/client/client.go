package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"goQuiz/internal"
	"goQuiz/internal/models"
)

type Client struct {
	BaseURL string
}

func NewClient() *Client {
	cfg := internal.LoadConfig()
	return &Client{BaseURL: cfg.BaseURL}
}

func (c *Client) GetQuestions() ([]models.Question, error) {
	resp, err := http.Get(c.BaseURL + "/questions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var questions []models.Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func (c *Client) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	data, err := json.Marshal(answers)
	if err != nil {
		return models.QuizResult{}, err
	}

	resp, err := http.Post(c.BaseURL+"/submit", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return models.QuizResult{}, err
	}
	defer resp.Body.Close()

	var result models.QuizResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.QuizResult{}, err
	}

	return result, nil
}
