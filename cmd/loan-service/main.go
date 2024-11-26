package main

import (
	"log"
	"net/http"

	"loan-service/configs"
	"loan-service/db"

	"loan-service/internal/handlers"
	loanHandler "loan-service/internal/handlers/loan"
	loanRepository "loan-service/internal/repositories/loan"
	loanService "loan-service/internal/services/loan"
)

func main() {
	config := configs.LoadConfig()

	// Initialize database
	err := db.InitializeDatabase(config.Database)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize repositories
	loanRepo := loanRepository.New(db.DB)

	// Initialize services
	loanService := loanService.New(loanRepo)

	// Initialize handlers
	loanHandler := loanHandler.New(loanService)
	handler := handlers.New(loanHandler)

	router := handlers.RegisterRoutes(&handler, db.DB)

	port := ":8080"
	log.Printf("Loan Service running on http://localhost%s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
