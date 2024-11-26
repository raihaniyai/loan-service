package main

import (
	"log"
	"net/http"

	"loan-service/configs"
	"loan-service/db"

	"loan-service/internal/handlers"
	actionHandler "loan-service/internal/handlers/action"
	loanHandler "loan-service/internal/handlers/loan"

	"loan-service/internal/repositories"
	actionRepository "loan-service/internal/repositories/action"
	investmentRepo "loan-service/internal/repositories/investment"
	loanRepository "loan-service/internal/repositories/loan"

	actionService "loan-service/internal/services/action"
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
	actionRepo := actionRepository.New(db.DB)
	dbRepo := repositories.New(db.DB)
	investmentRepo := investmentRepo.New(db.DB)
	loanRepo := loanRepository.New(db.DB)

	// Initialize services
	actionService := actionService.New(actionRepo, dbRepo, investmentRepo, loanRepo)
	loanService := loanService.New(actionRepo, dbRepo, loanRepo)

	// Initialize handlers
	actionHandler := actionHandler.New(actionService)
	loanHandler := loanHandler.New(loanService)
	handler := handlers.New(actionHandler, loanHandler)

	router := handlers.RegisterRoutes(&handler, db.DB)

	port := ":8080"
	log.Printf("Loan Service running on http://localhost%s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
