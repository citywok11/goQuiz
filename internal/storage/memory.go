package storage

import (
	"fmt"
	"sync"

	"goQuiz/internal/models"
)

type MemoryStorage struct {
	quiz    models.Quiz
	results []int
	mu      sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		quiz: models.Quiz{
			Questions: []models.Question{
				{
					ID:            1,
					Text:          "What is the capital of France?",
					Options:       []string{"London", "Berlin", "Paris", "Madrid"},
					CorrectAnswer: 2,
				},
				{
					ID:            2,
					Text:          "Which planet is known as the Red Planet?",
					Options:       []string{"Venus", "Mars", "Jupiter", "Saturn"},
					CorrectAnswer: 1,
				},
				// Add more questions here
			},
		},
		results: []int{},
	}
}

func (m *MemoryStorage) GetQuiz() models.Quiz {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.quiz
}

func (m *MemoryStorage) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	correctAnswers := 0
	for _, answer := range answers {
		for _, question := range m.quiz.Questions {
			if question.ID == answer.QuestionID && question.CorrectAnswer == answer.AnswerIndex {
				correctAnswers++
				break
			}
		}
	}

	// Calculate percentile before adding the new result
	var percentile float64
	if len(m.results) > 0 {
		betterScores := 0
		for _, score := range m.results {
			if score < correctAnswers {
				betterScores++
			}
		}
		percentile = float64(betterScores) / float64(len(m.results)) * 100
	} else {
		percentile = 100 // First submission is always in the 100th percentile
	}

	// Add the new result after calculating the percentile
	m.results = append(m.results, correctAnswers)

	result := models.QuizResult{
		CorrectAnswers: correctAnswers,
		TotalQuestions: len(m.quiz.Questions),
		Percentile:     percentile,
	}

	fmt.Printf("Submitted answers. Correct: %d, Total: %d, Percentile: %.2f\n",
		result.CorrectAnswers, result.TotalQuestions, result.Percentile)

	return result, nil
}
