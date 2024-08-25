package api

import (
	"encoding/json"
	"net/http"

	"goQuiz/internal/models"
	"goQuiz/internal/storage"
)

type Server struct {
	storage storage.QuizStorage
}

func NewServer(storage storage.QuizStorage) *Server {
	return &Server{storage: storage}
}

func (s *Server) GetQuestions(w http.ResponseWriter, r *http.Request) {
	quiz := s.storage.GetQuiz()
	json.NewEncoder(w).Encode(quiz.Questions)
}

func (s *Server) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var answers []models.UserAnswer
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.storage.SubmitAnswers(answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (s *Server) SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", s.GetQuestions)
	mux.HandleFunc("/submit", s.SubmitAnswers)
	return mux
}
