package main

import (
	"log"
	"net/http"
	"strconv"
	"warehouse/handlers"
	"warehouse/logic"
	"warehouse/routes"
	"warehouse/utils"
)

func main() {
	// Load config
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Error load config: %v", err)
	}

	// Initialize database connection using configuration
	err = utils.Connection(config)
	if err != nil {
		log.Fatalf("Error connection database: %v", err)
	}

	// Inisialisasi application logic
	authLogic := &logic.AuthLogic{JWTSecret: config.JWT.Secret, Config: config}
	authHandler := &handlers.AuthHandler{AuthLogic: authLogic}

	productLogic := &logic.ProductLogic{Config: config}
	productHandler := &handlers.ProductHandler{ProductLogic: productLogic}

	// Initialize handlers with logic
	handlers := &logic.Handlers{
		AuthHandler:    authHandler,
		ProductHandler: productHandler,
	}

	// Register routes and pass JWT secret for route security
	routes.RegisterRoutes(handlers, config.JWT.Secret)

	// Get port from configuration and start server
	port := config.Server.Port
	log.Printf("Server is running on port %d...", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
