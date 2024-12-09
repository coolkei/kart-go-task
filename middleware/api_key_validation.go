package middleware

import (
	"backend-challenge/models"
	"encoding/json"
	"net/http"
)

// ValidateAPIKey middleware to check the API key in request header
func ValidateAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api_key")

		// Replace with your actual API key
		validAPIKey := "YOUR_VALID_API_KEY"

		// Check if API key is provided and valid
		if apiKey == "" || apiKey != validAPIKey {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ApiResponse{
				Code:    401,
				Type:    "error",
				Message: "Unauthorized: Invalid or missing API key",
			})
			return
		}

		// Proceed to next handler if API key is valid
		next.ServeHTTP(w, r)
	})
}
