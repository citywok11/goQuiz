package storage

import (
	"errors"
	"fmt"
	"goQuiz/internal/models"
)

// SubmitAnswers processes the submitted answers and returns the number of correct answers.
// It returns an error if the number of submitted answers doesn't match the number of questions.
func (ms *MemoryStorage) SubmitAnswers(answers []models.UserAnswer) (int, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if len(answers) != len(ms.quiz.Questions) {
		return 0, errors.New("number of answers doesn't match number of questions")
	}

	correctAnswers := ms.calculateCorrectAnswers(answers)

	// Don't update results here, we'll do it in CalculateResult

	return correctAnswers, nil
}

// CalculateResult calculates the final quiz result, including the percentile.

func (ms *MemoryStorage) CalculateResult(correctAnswers int) models.QuizResult {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	percentile := ms.calculatePercentile(correctAnswers)

	// Update results after calculating percentile
	ms.updateResults(correctAnswers)

	return models.QuizResult{
		CorrectAnswers: correctAnswers,
		TotalQuestions: len(ms.quiz.Questions),
		Percentile:     percentile,
	}
}

func (ms *MemoryStorage) calculateCorrectAnswers(answers []models.UserAnswer) int {
	questionMap := make(map[int]models.Question)
	for _, q := range ms.quiz.Questions {
		questionMap[q.ID] = q
	}

	correctAnswers := 0
	for _, answer := range answers {
		if question, exists := questionMap[answer.QuestionID]; exists && question.CorrectAnswer == answer.AnswerIndex {
			correctAnswers++
		}
	}
	return correctAnswers
}

func (ms *MemoryStorage) updateResults(correctAnswers int) {
	ms.results = append(ms.results, correctAnswers)

	fmt.Println("ms reuslts appended", ms.results)

	if len(ms.results) > 1000 { // Limit the size of results
		ms.results = ms.results[1:]
	}
}

func (ms *MemoryStorage) calculatePercentile(correctAnswers int) float64 {
	totalResults := len(ms.results)

	if totalResults == 0 {
		fmt.Println("Debug: hit total results = 0")
		return 100 // First submission is always 100th percentile
	}

	count := 0
	for _, score := range ms.results {
		if score < correctAnswers {
			count++
		}
	}

	// Calculate the percentile
	// The current score is better than 'count' out of 'totalResults' previous scores
	return float64(count) / float64(totalResults) * 100
}
