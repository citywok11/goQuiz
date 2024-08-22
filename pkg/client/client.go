package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"goQuiz/internal/models"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
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
