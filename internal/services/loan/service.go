package loan

import (
	"loan-service/internal/repositories/loan"
)

type Service interface {
	CreateLoan(request CreateLoanRequest) (CreateLoanResult, error)
}

type service struct {
	loanRepository loan.Repository
}

func New(loanRepository loan.Repository) Service {
	return &service{
		loanRepository: loanRepository,
	}
}
