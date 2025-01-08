package logic

import (
	"net/http"
)

// AuthHandlerInterface defines the methods required for authentication-related actions.
type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

// ProductHandlerInterface defines the methods required for product-related actions.
type ProductHandlerInterface interface {
	AddProduct(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductByID(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

// Handlers groups the authentication and product handlers into a single struct.
type Handlers struct {
	AuthHandler    AuthHandlerInterface
	ProductHandler ProductHandlerInterface
}
