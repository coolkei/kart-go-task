package handlers

import (
	"backend-challenge/models"
	"encoding/json"
	"net/http"
	"strconv"
)

var products = []models.Product{
	{ID: "1", Name: "Waffle with Berries", Price: 6.50, Category: "Waffle"},
	{ID: "2", Name: "Vegan Bean Creme Brulee", Price: 7.0, Category: "Creme Brulee"},
	{ID: "3", Name: "Macaron Mix of Five", Price: 8.0, Category: "Macaron"},
	{ID: "4", Name: "Classic Tiramisu", Price: 5.50, Category: "Tiramisu"},
	{ID: "5", Name: "Pistachio Baklava", Price: 4.0, Category: "Baklava"},
	{ID: "6", Name: "Lemon Meringue Pie", Price: 5.0, Category: "Pie"},
	{ID: "7", Name: "Red Velvet Cake", Price: 4.50, Category: "Cake"},
	{ID: "8", Name: "Salted Caramel Brownie", Price: 5.50, Category: "Brownie"},
	{ID: "9", Name: "Vanilla panna Cotta", Price: 6.50, Category: "Panna Cotta"},
}

// ListProducts handles GET /product
func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProduct handles GET /product/{productID}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the productID from the URL
	productID := r.URL.Path[len("/product/"):]
	if productID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ApiResponse{
			Code:    400,
			Type:    "error",
			Message: "Invalid ID supplied",
		})
		return
	}

	// Validate that the productID is numeric
	if _, err := strconv.Atoi(productID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ApiResponse{
			Code:    400,
			Type:    "error",
			Message: "Invalid ID supplied",
		})
		return
	}

	// Search for the product in the product list
	for _, product := range products {
		if product.ID == productID {
			// Product found, return it with a 200 status
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	// Product not found, return 404
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.ApiResponse{
		Code:    404,
		Type:    "error",
		Message: "Product not found",
	})
}
