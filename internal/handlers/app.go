package handlers

import (
	"loan-service/internal/handlers/loan"
	"loan-service/internal/infrastructure/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	loanHandler loan.Handler
}

func New(loanHandler loan.Handler) Handler {
	return Handler{
		loanHandler: loanHandler,
	}
}

func RegisterRoutes(h *Handler, db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JWTMiddlewareWithDB(db))

	r.HandleFunc("/loans", h.loanHandler.CreateLoan).Methods("POST")

	return r
}
