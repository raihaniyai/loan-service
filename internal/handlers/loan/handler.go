package loan

import (
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
