package loan

import (
	"fmt"
	"log"
	"net/http"

	"loan-service/internal/services/loan"
)

type Handler struct {
	loanService loan.Service
}

func New(loanService loan.Service) Handler {
	return Handler{
		loanService: loanService,
	}
}

func (h *Handler) CreateLoan(w http.ResponseWriter, r *http.Request) {
	result, err := h.loanService.CreateLoan(loan.CreateLoanRequest{})
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", result)
}
