package action

import (
	"context"
	"loan-service/internal/infrastructure/nsq"
	"loan-service/internal/repositories"
	"loan-service/internal/repositories/action"
	"loan-service/internal/repositories/fund"
	"loan-service/internal/repositories/investment"
	"loan-service/internal/repositories/loan"
	"loan-service/internal/repositories/user"
)

type Service interface {
	UpdateLoan(ctx context.Context, request UpdateLoanRequest) (UpdateLoanResult, error)
	InvestLoan(ctx context.Context, request InvestLoanRequest) (InvestLoanResult, error)
	SendAgreementLetter(ctx context.Context, request SendAgreementLetterRequest) error
}

type service struct {
	actionRepository     action.Repository
	database             repositories.DB
	fundRepository       fund.Repository
	investmentRepository investment.Repository
	loanRepository       loan.Repository
	nsq                  *nsq.Publisher
	userRepository       user.Repository
}

func New(
	actionRepository action.Repository,
	database repositories.DB,
	fundRepository fund.Repository,
	investmentRepository investment.Repository,
	loanRepository loan.Repository,
	nsq *nsq.Publisher,
	userRepository user.Repository,
) Service {
	return &service{
		actionRepository:     actionRepository,
		database:             database,
		fundRepository:       fundRepository,
		investmentRepository: investmentRepository,
		loanRepository:       loanRepository,
		nsq:                  nsq,
		userRepository:       userRepository,
	}
}
