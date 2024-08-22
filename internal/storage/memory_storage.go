package storage

import (
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
