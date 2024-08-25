package quiz

import "goQuiz/internal/models"

type QuestionFetcher interface {
	FetchQuestions() ([]models.Question, error)
}

type AnswerCollector interface {
	CollectUserAnswers(questions []models.Question) []models.UserAnswer
}

type AnswerSubmitter interface {
	SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error)
}

type ResultDisplayer interface {
	DisplayResults(result models.QuizResult)
}
