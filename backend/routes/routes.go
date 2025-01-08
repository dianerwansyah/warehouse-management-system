package routes

import (
	"net/http"
	"strings"
	"warehouse/logic"
	"warehouse/middleware"
)

// RegisterRoutes registers routes for web application.
func RegisterRoutes(handlers *logic.Handlers, secret string) {
	// Route API, wrapped with CORS handler and calling login method from handler.
	http.Handle("/login", middleware.CORSHandler(http.HandlerFunc(handlers.AuthHandler.Login)))
	http.Handle("/register", middleware.CORSHandler(http.HandlerFunc(handlers.AuthHandler.Register)))

	http.Handle("/products/", middleware.CORSHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			if r.URL.Path == "/products/add" {
				handlers.ProductHandler.AddProduct(w, r)
			} else {
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		case http.MethodGet:
			if strings.HasPrefix(r.URL.Path, "/products/gets") {
				handlers.ProductHandler.GetProducts(w, r)
			} else if strings.HasPrefix(r.URL.Path, "/products/getbyid") {
				handlers.ProductHandler.GetProductByID(w, r)
			} else {
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		case http.MethodPut:
			if strings.HasPrefix(r.URL.Path, "/products/update") {
				handlers.ProductHandler.UpdateProduct(w, r)
			} else {
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		case http.MethodDelete:
			if strings.HasPrefix(r.URL.Path, "/products/delete/") {
				handlers.ProductHandler.DeleteProduct(w, r)
			} else {
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))
}
