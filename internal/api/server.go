package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

func (s *Server) ServeAPIDoc(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for: %s", r.URL.Path)
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.URL.Path == "/api-docs" {
		filePath := filepath.Join(currentDir, "api-docs.yaml")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Printf("api-docs.yaml file not found at %s", filePath)
			http.Error(w, "API documentation not found", http.StatusNotFound)
			return
		}
		log.Printf("Serving api-docs.yaml from: %s", filePath)
		w.Header().Set("Content-Type", "application/yaml")
		http.ServeFile(w, r, filePath)
	} else {
		filePath := filepath.Join(currentDir, "swagger-ui.html")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Printf("swagger-ui.html file not found at %s", filePath)
			http.Error(w, "Swagger UI not found", http.StatusNotFound)
			return
		}
		log.Printf("Serving swagger-ui.html from: %s", filePath)
		http.ServeFile(w, r, filePath)
	}
}

func (s *Server) SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", s.GetQuestions)
	mux.HandleFunc("/submit", s.SubmitAnswers)
	mux.HandleFunc("/api-docs", s.ServeAPIDoc)
	mux.HandleFunc("/docs", s.ServeAPIDoc)

	log.Println("Routes set up. API documentation available at /docs")

	return mux
}
