package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"warehouse/logic"
	"warehouse/models"
)

// ProductHandler holds logic for product-related operations
type ProductHandler struct {
	ProductLogic *logic.ProductLogic
}

// respondWithError sends an error response with a message and status code
func (ph *ProductHandler) respondWithError(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
}

// respondWithJSON sends a successful response with JSON payload
func (ph *ProductHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// AddProduct handles adding a new product
func (ph *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	// Decode request body into product model
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		ph.respondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Add product to database
	message, err := ph.ProductLogic.AddProduct(&product)
	if err != nil {
		ph.respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success message upon successful addition
	ph.respondWithJSON(w, http.StatusCreated, map[string]string{"message": message})
}

// GetProducts retrieves all products and returns them
func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.ProductLogic.GetProducts()
	if err != nil {
		ph.respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ph.respondWithJSON(w, http.StatusOK, products)
}

// GetProductByID retrieves a single product by its ID
func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/getbyid/")
	// Check if ID is empty
	if idStr == "" {
		ph.respondWithError(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// Convert string ID to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ph.respondWithError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Get product by ID
	product, err := ph.ProductLogic.GetProductByID(id)
	if err != nil {
		ph.respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if product == nil {
		ph.respondWithError(w, "Product not found", http.StatusNotFound)
		return
	}

	ph.respondWithJSON(w, http.StatusOK, product)
}

// UpdateProduct updates an existing product
func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	// Decode request body into product model
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		ph.respondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update product in database
	message, err := ph.ProductLogic.UpdateProduct(&product)
	if err != nil {
		ph.respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	ph.respondWithJSON(w, http.StatusCreated, map[string]string{"message": message})
}

// DeleteProduct deletes a product by its ID
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/delete/")
	// Check if ID is empty
	if idStr == "" {
		ph.respondWithError(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// Convert string ID to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ph.respondWithError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Delete product by ID
	message, err := ph.ProductLogic.DeleteProduct(id)
	if err != nil {
		ph.respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ph.respondWithJSON(w, http.StatusOK, map[string]string{"message": message})
}
