package storage

import (
	"goQuiz/internal/models"
)

func (m *MemoryStorage) GetQuiz() models.Quiz {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.quiz
}
