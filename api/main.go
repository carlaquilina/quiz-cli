package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/myapp/api/handlers"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define the endpoints
	router.HandleFunc("/questions", handlers.HandleQuestions).Methods("GET")
	router.HandleFunc("/answers", handlers.HandleAnswer).Methods("POST")
	router.HandleFunc("/stats", handlers.HandleStats).Methods("GET")

	// Serve the API
	log.Fatal(http.ListenAndServe(":8080", router))
}
