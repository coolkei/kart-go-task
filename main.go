package main

import (
	"backend-challenge/handlers"
	"log"
	"net/http"

	// "backend-challenge/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Product Routes
	r.HandleFunc("/product", handlers.ListProducts).Methods("GET")
	r.HandleFunc("/product/{productId}", handlers.GetProduct).Methods("GET")

	// Order Routes
	r.HandleFunc("/order", handlers.PlaceOrder).Methods("POST")

	// API Key validation middleware
	// http.Handle("/order", middleware.ValidateAPIKey(http.HandlerFunc(handlers.PlaceOrder)))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
