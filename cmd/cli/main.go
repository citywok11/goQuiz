package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"goQuiz/internal/models"

	"github.com/spf13/cobra"
)

const apiBaseURL = "http://localhost:8081"

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A simple quiz application",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new quiz",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Attempting to fetch questions...")
		questions, err := getQuestions()
		if err != nil {
			fmt.Println("Error fetching questions:", err)
			return
		}
		fmt.Printf("Successfully fetched %d questions\n", len(questions))

		var userAnswers []models.UserAnswer
		for _, q := range questions {
			fmt.Println(q.Text)
			for i, option := range q.Options {
				fmt.Printf("%d. %s\n", i+1, option)
			}

			var answer int
			fmt.Print("Enter your answer (1-4): ")
			fmt.Scan(&answer)

			userAnswers = append(userAnswers, models.UserAnswer{
				QuestionID:  q.ID,
				AnswerIndex: answer - 1,
			})
		}

		result, err := submitAnswers(userAnswers)
		if err != nil {
			fmt.Println("Error submitting answers:", err)
			return
		}

		fmt.Printf("You got %d out of %d questions correct.\n", result.CorrectAnswers, result.TotalQuestions)
		fmt.Printf("You performed better than %.2f%% of all quizzers.\n", result.Percentile)
	},
}

func getQuestions() ([]models.Question, error) {
	resp, err := http.Get(apiBaseURL + "/questions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var questions []models.Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func submitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	data, err := json.Marshal(answers)
	if err != nil {
		return models.QuizResult{}, err
	}

	resp, err := http.Post(apiBaseURL+"/submit", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return models.QuizResult{}, err
	}
	defer resp.Body.Close()

	var result models.QuizResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.QuizResult{}, err
	}

	return result, nil
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
