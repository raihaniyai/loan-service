package user

import (
	"context"

	"loan-service/internal/repositories"
	"loan-service/internal/repositories/fund"
	"loan-service/internal/repositories/user"
)

type Service interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (User, error)
	TopUpUserBalance(ctx context.Context, request TopUpUserBalanceRequest) (TopUpUserBalanceResult, error)
}

type service struct {
	database       repositories.DB
	fundRepository fund.Repository
	userRepository user.Repository
}

func New(database repositories.DB, fundRepository fund.Repository, userRepository user.Repository) Service {
	return &service{
		database:       database,
		fundRepository: fundRepository,
		userRepository: userRepository,
	}
}
