package storage

import "goQuiz/internal/models"

type QuizStorage interface {
	GetQuiz() models.Quiz
	SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error)
}
