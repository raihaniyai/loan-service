package action

import (
	"context"
	"loan-service/internal/repositories"
	"loan-service/internal/repositories/action"
	"loan-service/internal/repositories/fund"
	"loan-service/internal/repositories/investment"
	"loan-service/internal/repositories/loan"
)

type Service interface {
	UpdateLoan(ctx context.Context, request UpdateLoanRequest) (UpdateLoanResult, error)
	InvestLoan(ctx context.Context, request InvestLoanRequest) (InvestLoanResult, error)
}

type service struct {
	actionRepository     action.Repository
	database             repositories.DB
	fundRepository       fund.Repository
	investmentRepository investment.Repository
	loanRepository       loan.Repository
}

func New(
	actionRepository action.Repository,
	database repositories.DB,
	fundRepository fund.Repository,
	investmentRepository investment.Repository,
	loanRepository loan.Repository,
) Service {
	return &service{
		actionRepository:     actionRepository,
		database:             database,
		fundRepository:       fundRepository,
		investmentRepository: investmentRepository,
		loanRepository:       loanRepository,
	}
}
