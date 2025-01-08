package handlers

import (
	"encoding/json"
	"net/http"
	"warehouse/logic"
	"warehouse/models"
)

// AuthHandler holds reference to AuthLogic for handling user authentication
type AuthHandler struct {
	AuthLogic *logic.AuthLogic
}

// Register handles user registration
func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Register user in the database
	if err := ah.AuthLogic.Register(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return created user details as response
	response := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Login handles user authentication
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	// Decode request body to get login credentials
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate user and retrieve token
	token, authenticatedUser, err := ah.AuthLogic.Authenticate(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Return token and authenticated user details as response
	response := map[string]interface{}{
		"token":    token,
		"username": authenticatedUser.Username,
		"role":     authenticatedUser.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
