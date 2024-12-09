package handlers

import (
	"backend-challenge/models"
	"backend-challenge/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	orderIDCounter = 0
	mu             sync.Mutex
	files          = []string{"couponbase1.gz", "couponbase2.gz", "couponbase3.gz"}
)

// PlaceOrder handles POST /order
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq models.OrderReq
	err := json.NewDecoder(r.Body).Decode(&orderReq)
	if err != nil || len(orderReq.Items) == 0 {
		// Return 400 if the input is invalid
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ApiResponse{
			Code:    400,
			Type:    "error",
			Message: "Invalid input",
		})
		return
	}

	// Validate each item in the order
	for _, item := range orderReq.Items {
		// Check for missing productID
		if item.ProductID == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ApiResponse{
				Code:    400,
				Type:    "error",
				Message: "productId for item is required",
			})
			return
		}

		// Check for invalid or missing quantity
		if item.Quantity <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ApiResponse{
				Code:    400,
				Type:    "error",
				Message: "item quantity cannot be less than zero",
			})
			return
		}
	}

	// Validate promo code
	if orderReq.CouponCode != "" && !utils.IsValidPromoCode(orderReq.CouponCode, files) {
		// Return 422 if promo code validation fails
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.ApiResponse{
			Code:    422,
			Type:    "error",
			Message: "Invalid promo code",
		})
		return
	}

	// Validate product IDs in the order
	productQuantities := make(map[string]int)
	for _, item := range orderReq.Items {
		found := false
		for _, product := range products {
			if product.ID == item.ProductID {
				found = true
				productQuantities[item.ProductID] += item.Quantity
				break
			}
		}
		if !found {
			// Return 422 if an invalid product is specified
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(models.ApiResponse{
				Code:    422,
				Type:    "error",
				Message: fmt.Sprintf("Invalid product specified: %s", item.ProductID),
			})
			return
		}
	}

	// Generate order ID
	mu.Lock()
	orderIDCounter++
	orderID := orderIDCounter
	mu.Unlock()

	// Prepare the response
	var orderedItems []struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	}
	var orderedProducts []models.Product
	for productID, quantity := range productQuantities {
		for _, product := range products {
			if product.ID == productID {
				orderedItems = append(orderedItems, struct {
					ProductID string `json:"productId"`
					Quantity  int    `json:"quantity"`
				}{ProductID: productID, Quantity: quantity})
				orderedProducts = append(orderedProducts, product)
			}
		}
	}

	response := models.Order{
		ID:       formatOrderID(orderID),
		Items:    orderedItems,
		Products: orderedProducts,
	}

	// Return 200 with the order details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func formatOrderID(counter int) string {
	return "0000-0000-0000-" + formatCounter(counter)
}

func formatCounter(counter int) string {
	return fmt.Sprintf("%04d", counter)
}
