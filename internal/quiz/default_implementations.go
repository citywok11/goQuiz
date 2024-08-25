package quiz

import (
	"fmt"
	"goQuiz/internal/api"
	"goQuiz/internal/models"
)

type defaultQuestionFetcher struct{}

func (d defaultQuestionFetcher) FetchQuestions() ([]models.Question, error) {
	return api.GetQuestions()
}

type defaultAnswerCollector struct{}

func (d defaultAnswerCollector) CollectUserAnswers(questions []models.Question) []models.UserAnswer {
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

type defaultAnswerSubmitter struct{}

func (d defaultAnswerSubmitter) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	return api.SubmitAnswers(answers)
}

type defaultResultDisplayer struct{}

func (d defaultResultDisplayer) DisplayResults(result models.QuizResult) {
	fmt.Printf("You got %d out of %d questions correct.\n", result.CorrectAnswers, result.TotalQuestions)
	fmt.Printf("You performed better than %.2f%% of all quizzers.\n", result.Percentile)
}
