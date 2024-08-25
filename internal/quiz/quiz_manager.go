package quiz

import "goQuiz/internal/models"

type QuizManager struct {
	fetcher   QuestionFetcher
	collector AnswerCollector
	submitter AnswerSubmitter
	displayer ResultDisplayer
}

func NewQuizManager() *QuizManager {
	return &QuizManager{
		fetcher:   defaultQuestionFetcher{},
		collector: defaultAnswerCollector{},
		submitter: defaultAnswerSubmitter{},
		displayer: defaultResultDisplayer{},
	}
}

func (qm *QuizManager) FetchQuestions() ([]models.Question, error) {
	return qm.fetcher.FetchQuestions()
}

func (qm *QuizManager) CollectUserAnswers(questions []models.Question) []models.UserAnswer {
	return qm.collector.CollectUserAnswers(questions)
}

func (qm *QuizManager) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	return qm.submitter.SubmitAnswers(answers)
}

func (qm *QuizManager) DisplayResults(result models.QuizResult) {
	qm.displayer.DisplayResults(result)
}
