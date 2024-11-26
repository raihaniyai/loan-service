package action

import (
	"context"
	"loan-service/internal/repositories"
	"loan-service/internal/repositories/action"
	"loan-service/internal/repositories/investment"
	"loan-service/internal/repositories/loan"
)

type Service interface {
	ApproveLoan(ctx context.Context, request ApproveLoanRequest) (ApproveLoanResult, error)
	InvestLoan(ctx context.Context, request InvestLoanRequest) (InvestLoanResult, error)
}

type service struct {
	actionRepository     action.Repository
	database             repositories.DB
	investmentRepository investment.Repository
	loanRepository       loan.Repository
}

func New(actionRepository action.Repository, database repositories.DB, investmentRepository investment.Repository, loanRepository loan.Repository) Service {
	return &service{
		actionRepository:     actionRepository,
		database:             database,
		investmentRepository: investmentRepository,
		loanRepository:       loanRepository,
	}
}
