package handlers

import (
	"loan-service/internal/handlers/action"
	"loan-service/internal/handlers/loan"
	"loan-service/internal/infrastructure/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	actionHandler action.Handler
	loanHandler   loan.Handler
}

func New(actionHandler action.Handler, loanHandler loan.Handler) Handler {
	return Handler{
		actionHandler: actionHandler,
		loanHandler:   loanHandler,
	}
}

func RegisterRoutes(h *Handler, db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JWTMiddlewareWithDB(db))

	r.HandleFunc("/loans", h.loanHandler.CreateLoan).Methods("POST")
	r.HandleFunc("/loans/{loanID}/approve", h.actionHandler.ApproveLoan).Methods("POST")
	r.HandleFunc("/loans/{loanID}/disburse", h.actionHandler.DisburseLoan).Methods("POST")
	r.HandleFunc("/loans/{loanID}/invest", h.actionHandler.InvestLoan).Methods("POST")

	// TODO
	r.HandleFunc("/loans", h.loanHandler.GetLoans).Methods("GET")
	r.HandleFunc("/loans/{loanID}", h.loanHandler.GetLoanDetails).Methods("GET")

	return r
}
