package main

import (
	"fmt"
	"log"
	"net/http"

	"goQuiz/internal"
	"goQuiz/internal/api"
	"goQuiz/internal/storage"
)

func setupServer(storage storage.QuizStorage, serverAddr string) *http.Server {
	server := api.NewServer(storage)
	handler := server.SetupRoutes()
	return &http.Server{
		Addr:    serverAddr,
		Handler: handler,
	}
}

func main() {
	cfg := internal.LoadConfig()
	memoryStorage := storage.NewMemoryStorage()
	serverAddr := fmt.Sprintf(":%d", cfg.ServerPort)
	server := setupServer(memoryStorage, serverAddr)

	fmt.Printf("Starting server on %s\n", serverAddr)
	log.Fatal(server.ListenAndServe())
}
