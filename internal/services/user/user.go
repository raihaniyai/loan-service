package user

import (
	"context"
	"errors"
	"log"

	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/constant"
)

func (svc *service) CreateUser(ctx context.Context, request CreateUserRequest) (User, error) {
	var err error

	// get user by email
	user, err := svc.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return User{}, err
	}

	if user != nil {
		return User{}, errors.New("user already exists")
	}

	tx := svc.database.BeginTx()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	userID, err := svc.userRepository.SetUser(ctx, tx, &entity.User{
		Name:        request.Name,
		Role:        request.Role,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	})
	if err != nil {
		return User{}, err
	}

	if request.Role != constant.UserRoleAdmin {
		_, err = svc.fundRepository.SetFund(ctx, tx, &entity.Fund{
			UserID: userID,
		})
		if err != nil {
			return User{}, err
		}
	}

	errCommit := tx.Commit().Error
	if errCommit != nil {
		return User{}, errCommit
	}

	return User{
		UserID:      userID,
		Name:        request.Name,
		Role:        request.Role,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}, nil
}

func (svc *service) TopUpUserBalance(ctx context.Context, request TopUpUserBalanceRequest) (TopUpUserBalanceResult, error) {
	if request.UserRole != constant.UserRoleBorrower && request.UserRole != constant.UserRoleInvestor {
		log.Println("SVC.TUUB00 | [TopUpUserBalance] User is not eligible to top up balance")
		return TopUpUserBalanceResult{}, errors.New("user is not eligible to top up balance")
	}

	currentBalance, err := svc.fundRepository.GetBalanceByUserID(ctx, request.UserID)
	if err != nil {
		log.Println("SVC.TUUB00 | [TopUpUserBalance] Error getting balance by user id:", err)
		return TopUpUserBalanceResult{}, err
	}

	err = svc.fundRepository.UpdateBalanceByUserID(ctx, nil, request.UserID, currentBalance+request.TopUpAmount)
	if err != nil {
		log.Println("SVC.TUUB01 | [TopUpUserBalance] Error updating balance by user id:", err)
		return TopUpUserBalanceResult{}, err
	}

	return TopUpUserBalanceResult{
		TotalBalanceAmount: currentBalance + request.TopUpAmount,
	}, nil
}
