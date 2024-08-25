package main

import (
	"fmt"
	"log"
	"net/http"

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
	memoryStorage := storage.NewMemoryStorage()
	serverAddr := ":8080"
	server := setupServer(memoryStorage, serverAddr)

	fmt.Printf("Starting server on %s\n", serverAddr)
	log.Fatal(server.ListenAndServe())
}
