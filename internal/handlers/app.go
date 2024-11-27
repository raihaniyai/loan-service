package handlers

import (
	"loan-service/internal/handlers/action"
	"loan-service/internal/handlers/loan"
	"loan-service/internal/handlers/user"
	"loan-service/internal/infrastructure/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	actionHandler action.Handler
	loanHandler   loan.Handler
	userHandler   user.Handler
}

func New(actionHandler action.Handler, loanHandler loan.Handler, userHandler user.Handler) Handler {
	return Handler{
		actionHandler: actionHandler,
		loanHandler:   loanHandler,
		userHandler:   userHandler,
	}
}

func RegisterRoutes(h *Handler, db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JWTMiddlewareWithDB(db))

	r.HandleFunc("/loans", h.loanHandler.CreateLoan).Methods("POST")
	r.HandleFunc("/loans/{loanID}/approve", h.actionHandler.ApproveLoan).Methods("POST")
	r.HandleFunc("/loans/{loanID}/disburse", h.actionHandler.DisburseLoan).Methods("POST")
	r.HandleFunc("/loans/{loanID}/invest", h.actionHandler.InvestLoan).Methods("POST")

	// opt routes
	r.HandleFunc("/funds/topup", h.userHandler.TopUpUserBalance).Methods("POST")
	r.HandleFunc("/users", h.userHandler.CreateUser).Methods("POST")
	// r.HandleFunc("/loans", h.loanHandler.GetLoans).Methods("GET")
	// r.HandleFunc("/loans/{loanID}", h.loanHandler.GetLoanDetails).Methods("GET")

	return r
}
