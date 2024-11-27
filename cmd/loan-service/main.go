package main

import (
	"log"
	"net/http"

	"loan-service/configs"
	"loan-service/db"

	"loan-service/internal/handlers"
	actionHandler "loan-service/internal/handlers/action"
	loanHandler "loan-service/internal/handlers/loan"
	userHandler "loan-service/internal/handlers/user"

	"loan-service/internal/repositories"
	actionRepository "loan-service/internal/repositories/action"
	fundRepository "loan-service/internal/repositories/fund"
	investmentRepo "loan-service/internal/repositories/investment"
	loanRepository "loan-service/internal/repositories/loan"
	userRepository "loan-service/internal/repositories/user"

	actionService "loan-service/internal/services/action"
	loanService "loan-service/internal/services/loan"
	userService "loan-service/internal/services/user"

	"loan-service/internal/infrastructure/nsq"
)

func main() {
	config := configs.LoadConfig()

	// Initialize database
	err := db.InitializeDatabase(config.Database)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize NSQ Publisher
	nsqPublisher, err := nsq.NewPublisher(config.NSQ.NSQDAddress)
	if err != nil {
		log.Fatalf("Error initializing NSQ producer: %v", err)
	}
	defer nsqPublisher.Stop()

	// Initialize NSQ Consumer
	nsqMessageHandler := &handlers.NSQMessageHandler{}
	nsqConsumer, err := nsq.NewConsumer(
		config.NSQ.Topic,
		config.NSQ.Channel,
		config.NSQ.LookupDAddress,
		nsqMessageHandler,
	)
	if err != nil {
		log.Fatalf("Error initializing NSQ consumer: %v", err)
	}
	defer nsqConsumer.Stop()

	// Initialize repositories
	actionRepo := actionRepository.New(db.DB)
	dbRepo := repositories.New(db.DB)
	fundRepo := fundRepository.New(db.DB)
	investmentRepo := investmentRepo.New(db.DB)
	loanRepo := loanRepository.New(db.DB)
	userRepo := userRepository.New(db.DB)

	// Initialize services
	actionService := actionService.New(actionRepo, dbRepo, fundRepo, investmentRepo, nsqPublisher, loanRepo)
	loanService := loanService.New(actionRepo, dbRepo, loanRepo)
	userService := userService.New(dbRepo, fundRepo, userRepo)

	// Initialize handlers
	actionHandler := actionHandler.New(actionService)
	loanHandler := loanHandler.New(loanService)
	userHandler := userHandler.New(userService)
	handler := handlers.New(actionHandler, loanHandler, userHandler)

	router := handlers.RegisterRoutes(&handler, db.DB)

	port := ":8080"
	log.Printf("⚡ Loan Service running on http://localhost%s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
