package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uyendava/go-calculator/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/calculate", handlers.CalculateHandler)

	fmt.Println("Server is running on http://localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}
}
