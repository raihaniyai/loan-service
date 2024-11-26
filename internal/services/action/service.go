package action

import (
	"context"
	"loan-service/internal/repositories"
	"loan-service/internal/repositories/action"
	"loan-service/internal/repositories/loan"
)

type Service interface {
	ApproveLoan(ctx context.Context, request ApproveLoanRequest) (ApproveLoanResult, error)
}

type service struct {
	actionRepository action.Repository
	database         repositories.DB
	loanRepository   loan.Repository
}

func New(actionRepository action.Repository, database repositories.DB, loanRepository loan.Repository) Service {
	return &service{
		actionRepository: actionRepository,
		database:         database,
		loanRepository:   loanRepository,
	}
}
