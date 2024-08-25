package quiz

import (
	"bufio"
	"fmt"
	"goQuiz/internal/api"
	"goQuiz/internal/models"
	"os"
	"strconv"
	"strings"
)

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

func displayQuestion(q models.Question) {
	fmt.Println(q.Text)
	for i, option := range q.Options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

func promptForAnswer(minOption, maxOption int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter your answer (%d-%d): ", minOption, maxOption)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		answer, err := strconv.Atoi(input)
		if err != nil || answer < minOption || answer > maxOption {
			fmt.Printf("Invalid input. Please enter a number between %d and %d.\n", minOption, maxOption)
			continue
		}
		return answer
	}
}
