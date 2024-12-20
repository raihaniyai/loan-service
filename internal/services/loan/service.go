package loan

import (
	"context"

	"loan-service/internal/repositories"
	"loan-service/internal/repositories/action"
	"loan-service/internal/repositories/loan"
)

type Service interface {
	CreateLoan(ctx context.Context, request CreateLoanRequest) (CreateLoanResult, error)
}

type service struct {
	database         repositories.DB
	actionRepository action.Repository
	loanRepository   loan.Repository
}

func New(actionRepository action.Repository, database repositories.DB, loanRepository loan.Repository) Service {
	return &service{
		actionRepository: actionRepository,
		database:         database,
		loanRepository:   loanRepository,
	}
}
