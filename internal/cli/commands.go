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

func NewQuizCommand(
	fetcher quiz.QuestionFetcher,
	collector quiz.AnswerCollector,
	submitter quiz.AnswerSubmitter,
	displayer quiz.ResultDisplayer,
) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start a new quiz",
		Run: func(cmd *cobra.Command, args []string) {
			runQuiz(fetcher, collector, submitter, displayer)
		},
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	startCmd := NewQuizCommand(
		&quiz.DefaultQuestionFetcher{},
		&quiz.DefaultAnswerCollector{},
		&quiz.DefaultAnswerSubmitter{},
		&quiz.DefaultResultDisplayer{},
	)
	rootCmd.AddCommand(startCmd)
}

func runQuiz(
	fetcher quiz.QuestionFetcher,
	collector quiz.AnswerCollector,
	submitter quiz.AnswerSubmitter,
	displayer quiz.ResultDisplayer,
) {
	quizManager := quiz.NewQuizManager(fetcher, collector, submitter, displayer)

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
