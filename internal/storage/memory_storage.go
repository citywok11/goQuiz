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
				{
					ID:            3,
					Text:          "Who wrote 'Romeo and Juliet'?",
					Options:       []string{"William Shakespeare", "Mark Twain", "Charles Dickens", "Jane Austen"},
					CorrectAnswer: 0,
				},
				{
					ID:            4,
					Text:          "What is the chemical symbol for water?",
					Options:       []string{"O2", "H2O", "CO2", "NaCl"},
					CorrectAnswer: 1,
				},
				{
					ID:            5,
					Text:          "Which continent is the Sahara Desert located on?",
					Options:       []string{"Asia", "Australia", "Africa", "South America"},
					CorrectAnswer: 2,
				},
			},
		},
		results: []int{},
	}
}

var _ QuizStorage = (*MemoryStorage)(nil)

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
		percentile = 100
	}

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
