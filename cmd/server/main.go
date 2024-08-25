package main

import (
	"fmt"
	"log"
	"net/http"

	"goQuiz/internal/api"
	"goQuiz/internal/storage"
)

func setupServer(storage storage.QuizStorage) *http.Server {
	server := api.NewServer(storage)
	handler := server.SetupRoutes()
	return &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}
}

func main() {
	memoryStorage := storage.NewMemoryStorage()
	server := setupServer(memoryStorage)

	fmt.Println("Starting server on :8081")
	log.Fatal(server.ListenAndServe())
}
