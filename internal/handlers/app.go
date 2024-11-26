package handlers

import (
	"loan-service/internal/handlers/loan"

	"github.com/gorilla/mux"
)

type Handler struct {
	loanHandler loan.Handler
}

func New(loanHandler loan.Handler) Handler {
	return Handler{
		loanHandler: loanHandler,
	}
}

func RegisterRoutes(h *Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/loans", h.loanHandler.CreateLoan).Methods("POST")

	return r
}
