package storage

import (
	"goQuiz/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryStorage_SubmitAnswers_Percentile(t *testing.T) {
	storage := NewMemoryStorage()

	// First submission: 2 correct answers out of 2
	result1, _ := storage.SubmitAnswers([]models.UserAnswer{
		{QuestionID: 1, AnswerIndex: 2}, // Correct
		{QuestionID: 2, AnswerIndex: 1}, // Correct
	})
	assert.Equal(t, 100.0, result1.Percentile)

	// Second submission: 1 correct answer out of 2
	result2, _ := storage.SubmitAnswers([]models.UserAnswer{
		{QuestionID: 1, AnswerIndex: 2}, // Correct
		{QuestionID: 2, AnswerIndex: 0}, // Incorrect
	})
	assert.Equal(t, 0.0, result2.Percentile)

	// Third submission: 2 correct answers out of 2
	result3, _ := storage.SubmitAnswers([]models.UserAnswer{
		{QuestionID: 1, AnswerIndex: 2}, // Correct
		{QuestionID: 2, AnswerIndex: 1}, // Correct
	})
	assert.Equal(t, 50.0, result3.Percentile)
}

func TestMemoryStorage_SubmitAnswers_FirstSubmissionPercentileAlways100(t *testing.T) {
	storage := NewMemoryStorage()

	// First submission: 0 correct answers out of 2
	result1, _ := storage.SubmitAnswers([]models.UserAnswer{
		{QuestionID: 1, AnswerIndex: 3}, // Incorrect
		{QuestionID: 2, AnswerIndex: 2}, // Incorrect
	})
	assert.Equal(t, 100.0, result1.Percentile)
}
