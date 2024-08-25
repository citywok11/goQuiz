package cli

import (
	"fmt"
	"goQuiz/internal/quiz"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A simple quiz application",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new quiz",
	Run:   runQuiz,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runQuiz(cmd *cobra.Command, args []string) {
	quizManager := quiz.NewQuizManager(
		&quiz.DefaultQuestionFetcher{},
		&quiz.DefaultAnswerCollector{},
		&quiz.DefaultAnswerSubmitter{},
		&quiz.DefaultResultDisplayer{},
	)

	questions, err := quizManager.FetchQuestions()
	if err != nil {
		fmt.Println("Error fetching questions:", err)
		return
	}

	userAnswers := quizManager.CollectUserAnswers(questions)

	result, err := quizManager.SubmitAnswers(userAnswers)
	if err != nil {
		fmt.Println("Error submitting answers:", err)
		return
	}

	quizManager.DisplayResults(result)
}
