package client

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var questions []models.Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return questions, nil
}

func (c *Client) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	data, err := json.Marshal(answers)
	if err != nil {
		return models.QuizResult{}, fmt.Errorf("failed to marshal answers: %w", err)
	}

	resp, err := http.Post(c.BaseURL+"/submit", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return models.QuizResult{}, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.QuizResult{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result models.QuizResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.QuizResult{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}
