package api

import (
	"bytes"
	"encoding/json"
	"goQuiz/internal/models"
	"net/http"
)

const apiBaseURL = "http://localhost:8081"

func GetQuestions() ([]models.Question, error) {
	resp, err := http.Get(apiBaseURL + "/questions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var questions []models.Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	return questions, err
}

func SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	data, err := json.Marshal(answers)
	if err != nil {
		return models.QuizResult{}, err
	}

	resp, err := http.Post(apiBaseURL+"/submit", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return models.QuizResult{}, err
	}
	defer resp.Body.Close()

	var result models.QuizResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
