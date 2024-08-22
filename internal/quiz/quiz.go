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

func FetchQuestions() ([]models.Question, error) {
	return api.GetQuestions()
}

func CollectUserAnswers(questions []models.Question) []models.UserAnswer {
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

func SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	return api.SubmitAnswers(answers)
}

func DisplayResults(result models.QuizResult) {
	fmt.Printf("You got %d out of %d questions correct.\n", result.CorrectAnswers, result.TotalQuestions)
	fmt.Printf("You performed better than %.2f%% of all quizzers.\n", result.Percentile)
}
