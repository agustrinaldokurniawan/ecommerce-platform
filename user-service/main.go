package main

import (
	"log"
	"net/http"

	"example.com/m/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Println("User service is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
