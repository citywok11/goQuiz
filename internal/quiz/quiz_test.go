package quiz

import (
	"fmt"
	"goQuiz/internal/api"
	"goQuiz/internal/models"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	testServer *httptest.Server
	serverOnce sync.Once
)

func setupTestServer() *httptest.Server {
	serverOnce.Do(func() {
		testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond) // Simulate a short delay
			switch r.URL.Path {
			case "/questions":
				w.Write([]byte(`[{"id":1,"text":"Test Question"}]`))
			case "/submit":
				w.Write([]byte(`{"correctAnswers": 1, "totalQuestions": 1, "percentile": 75.0}`))
			default:
				http.NotFound(w, r)
			}
		}))
	})
	return testServer
}

func TestMain(m *testing.M) {
	fmt.Println("Starting test setup")

	// Setup
	server := setupTestServer()
	originalURL := api.BaseURL
	api.BaseURL = server.URL

	defer func() {
		fmt.Println("Starting test teardown")
		server.Close()
		api.BaseURL = originalURL
		fmt.Println("Tests completed, exiting")
	}()

	fmt.Println("Running tests")

	// Run tests with a timeout
	done := make(chan int, 1)
	go func() {
		done <- m.Run()
	}()

	var code int
	select {
	case code = <-done:
		fmt.Println("Tests finished normally")
	case <-time.After(10 * time.Second):
		fmt.Println("Tests timed out, possible goroutine leak")
		code = 1
	}

	os.Exit(code)
}

type mockQuestionFetcher struct {
	mock.Mock
}

func (m *mockQuestionFetcher) FetchQuestions() ([]models.Question, error) {
	args := m.Called()
	return args.Get(0).([]models.Question), args.Error(1)
}

type mockAnswerCollector struct {
	mock.Mock
}

func (m *mockAnswerCollector) CollectUserAnswers(questions []models.Question) []models.UserAnswer {
	args := m.Called(questions)
	return args.Get(0).([]models.UserAnswer)
}

type mockAnswerSubmitter struct {
	mock.Mock
}

func (m *mockAnswerSubmitter) SubmitAnswers(answers []models.UserAnswer) (models.QuizResult, error) {
	args := m.Called(answers)
	return args.Get(0).(models.QuizResult), args.Error(1)
}

type mockResultDisplayer struct {
	mock.Mock
}

func (m *mockResultDisplayer) DisplayResults(result models.QuizResult) {
	m.Called(result)
}

func TestQuizManager_FetchQuestions(t *testing.T) {
	t.Parallel()
	t.Log("Starting TestQuizManager_FetchQuestions")

	mockFetcher := new(mockQuestionFetcher)
	quizManager := &QuizManager{fetcher: mockFetcher}

	expectedQuestions := []models.Question{{ID: 1, Text: "Test Question"}}
	mockFetcher.On("FetchQuestions").Return(expectedQuestions, nil)

	questions, err := quizManager.FetchQuestions()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(questions, expectedQuestions) {
		t.Errorf("Expected questions %+v, got %+v", expectedQuestions, questions)
	}

	mockFetcher.AssertExpectations(t)
	t.Log("TestQuizManager_FetchQuestions completed successfully")
}

func TestQuizManager_CollectUserAnswers(t *testing.T) {
	t.Parallel()
	mockCollector := new(mockAnswerCollector)
	quizManager := &QuizManager{collector: mockCollector}

	questions := []models.Question{{ID: 1, Text: "Test Question"}}
	expectedAnswers := []models.UserAnswer{{QuestionID: 1, AnswerIndex: 0}}
	mockCollector.On("CollectUserAnswers", questions).Return(expectedAnswers)

	answers := quizManager.CollectUserAnswers(questions)

	assert.Equal(t, expectedAnswers, answers)
	mockCollector.AssertExpectations(t)
	t.Log("TestQuizManager_CollectUserAnswers completed successfully")
}

func TestQuizManager_SubmitAnswers(t *testing.T) {
	t.Parallel()
	mockSubmitter := new(mockAnswerSubmitter)
	quizManager := &QuizManager{submitter: mockSubmitter}

	answers := []models.UserAnswer{{QuestionID: 1, AnswerIndex: 0}}
	expectedResult := models.QuizResult{CorrectAnswers: 1, TotalQuestions: 1, Percentile: 75.0}
	mockSubmitter.On("SubmitAnswers", answers).Return(expectedResult, nil)

	result, err := quizManager.SubmitAnswers(answers)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	mockSubmitter.AssertExpectations(t)
	t.Log("TestQuizManager_SubmitAnswers completed successfully")
}

func TestQuizManager_DisplayResults(t *testing.T) {
	t.Parallel()
	mockDisplayer := new(mockResultDisplayer)
	quizManager := &QuizManager{displayer: mockDisplayer}

	result := models.QuizResult{CorrectAnswers: 1, TotalQuestions: 1, Percentile: 75.0}
	mockDisplayer.On("DisplayResults", result).Return()

	quizManager.DisplayResults(result)

	mockDisplayer.AssertExpectations(t)
	t.Log("TestQuizManager_DisplayResults completed successfully")
}
