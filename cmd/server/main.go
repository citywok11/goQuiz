package main

import (
	"fmt"
	"log"
	"net/http"

	"goQuiz/internal/api"
	"goQuiz/internal/storage"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	server := api.NewServer(memoryStorage)

	handler := server.SetupRoutes()

	fmt.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
