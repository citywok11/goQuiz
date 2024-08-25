package quiz

import (
	"fmt"
	"goQuiz/internal/api"
	"goQuiz/internal/models"
)

// DefaultQuestionFetcher implements QuestionFetcher
type DefaultQuestionFetcher struct{}

func (d DefaultQuestionFetcher) FetchQuestions() ([]models.Question, error) {
	return api.GetQuestions()
}

// DefaultAnswerCollector implements AnswerCollector
type DefaultAnswerCollector struct{}

func (d DefaultAnswerCollector) CollectUserAnswers(questions []models.Question) []models.UserAnswer {
	var userAnswers []models.UserAnswer
	for _, q := range questions {
		displayQuestion(q)
		answer := promptForAnswer(1, len(q.Options))
		userAnswers = append(userAnswers, models.UserAnswer{
			QuestionID:  q.ID,
			AnswerIndex: answer - 1,
		})
	}
	return userAnswers
}

// DefaultAnswerSubmitter implements AnswerSubmitter
type DefaultAnswerSubmitter struct{}

func (d DefaultAnswerSubmitter) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	return api.SubmitAnswers(answers)
}

// DefaultResultDisplayer implements ResultDisplayer
type DefaultResultDisplayer struct{}

func (d DefaultResultDisplayer) DisplayResults(result models.QuizResult) {
	fmt.Printf("You got %d out of %d questions correct.\n", result.CorrectAnswers, result.TotalQuestions)
	fmt.Printf("You performed better than %.2f%% of all quizzers.\n", result.Percentile)
}
