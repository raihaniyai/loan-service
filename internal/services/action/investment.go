package action

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/constant"
)

func (svc *service) InvestLoan(ctx context.Context, request InvestLoanRequest) (InvestLoanResult, error) {
	if request.UserRole != constant.UserRoleInvestor {
		log.Println("SVC.IL00 | [InvestLoan] User is not an investor")
		return InvestLoanResult{}, errors.New("user is not an investor")
	}

	loan, err := svc.loanRepository.GetLoanByID(ctx, request.LoanID)
	if err != nil {
		log.Println("SVC.IL01 | [InvestLoan] Error getting loan:", err)
		return InvestLoanResult{}, err
	}

	if loan == nil {
		log.Println("SVC.IL02 | [InvestLoan] Loan not found")
		return InvestLoanResult{}, errors.New("loan not found")
	}

	if loan.Status != constant.LoanStatusApproved {
		log.Println("SVC.IL03 | [InvestLoan] Loan is not eligible to be invested")
		return InvestLoanResult{}, errors.New("loan is not eligible to be invested")
	}

	userBalance, err := svc.fundRepository.GetBalanceByUserID(ctx, request.UserID)
	if err != nil {
		log.Println("SVC.IL15 | [InvestLoan] Error getting user balance:", err)
		return InvestLoanResult{}, err
	}

	if userBalance < request.InvestmentAmount {
		log.Println("SVC.IL16 | [InvestLoan] The balance is not enough")
		return InvestLoanResult{}, errors.New("your balance is not enough")
	}

	if request.InvestmentAmount > loan.PrincipalAmount {
		log.Println("SVC.IL04 | [InvestLoan] Investment amount is too high")
		return InvestLoanResult{}, errors.New("investment amount is too high")
	}

	totalCurrentInvestmentAmount, err := svc.investmentRepository.GetTotalInvestmentAmountByLoanID(ctx, request.LoanID)
	if err != nil {
		log.Println("SVC.IL05 | [InvestLoan] Error getting total investment amount:", err)
		return InvestLoanResult{}, err
	}
	totalInvestmentAmount := totalCurrentInvestmentAmount + request.InvestmentAmount
	if totalInvestmentAmount > loan.PrincipalAmount {
		log.Println("SVC.IL06 | [InvestLoan] Investment amount is too high")
		return InvestLoanResult{}, errors.New("investment amount is too high")
	}

	investment, err := svc.investmentRepository.GetInvestmentByLoanIDAndInvestorID(ctx, request.LoanID, request.UserID)
	if err != nil {
		log.Println("SVC.IL07 | [InvestLoan] Error getting investment:", err)
		return InvestLoanResult{}, err
	}

	if investment != nil {
		log.Println("SVC.IL08 | [InvestLoan] User has already invested in this loan")
		return InvestLoanResult{}, errors.New("user has already invested in this loan")
	}

	tx := svc.database.BeginTx()
	defer func() {
		if err != nil {
			err = svc.database.Rollback(tx)
			if err != nil {
				log.Println("SVC.IL09 | [InvestLoan] Error rolling back transaction:", err)
			}
		}
	}()

	currentTime := time.Now()
	investment = &entity.Investment{
		LoanID:           request.LoanID,
		InvestorID:       request.UserID,
		InvestmentAmount: request.InvestmentAmount,
		CreatedAt:        currentTime,
	}

	investment.InvestmentID, err = svc.investmentRepository.SetInvestment(ctx, tx, investment)
	if err != nil {
		log.Println("SVC.IL10 | [InvestLoan] Error inserting investment:", err)
		return InvestLoanResult{}, err
	}

	if totalInvestmentAmount == loan.PrincipalAmount {
		loan.Status = constant.LoanStatusInvested
		loan.UpdatedAt = currentTime
		loan.UpdatedBy = request.UserID

		err = svc.loanRepository.UpdateLoan(ctx, tx, loan)
		if err != nil {
			log.Println("SVC.IL11 | [InvestLoan] Error updating loan:", err)
			return InvestLoanResult{}, err
		}

		//set action
		action := &entity.Action{
			LoanID:     loan.LoanID,
			ActionType: constant.ActionTypeInvestLoan,
			CreatedBy:  0,
			CreatedAt:  currentTime,
		}
		action.ActionID, err = svc.actionRepository.SetAction(ctx, tx, action)
		if err != nil {
			log.Println("SVC.IL12 | [InvestLoan] Error inserting action:", err)
			return InvestLoanResult{}, err
		}

		investments, err := svc.investmentRepository.GetInvestmentsByLoanID(ctx, request.LoanID)
		if err != nil {
			log.Println("SVC.IL13 | [InvestLoan] Error getting investments:", err)
			return InvestLoanResult{}, err
		}

		for _, investment := range investments {
			// TODO: publish kafka message
			// send pdf documents to all investors (publish kafka message)
			fmt.Println(investment.InvestorID)
		}
	}

	err = svc.fundRepository.UpdateBalanceByUserID(ctx, tx, request.UserID, userBalance-request.InvestmentAmount)
	if err != nil {
		log.Println("SVC.IL13 | [InvestLoan] Error updating balance:", err)
		return InvestLoanResult{}, err
	}

	errCommit := svc.database.Commit(tx)
	if errCommit != nil {
		log.Println("SVC.IL14 | [InvestLoan] Error committing transaction:", errCommit)
		return InvestLoanResult{}, errCommit
	}

	return InvestLoanResult{
		InvestmentID: investment.InvestmentID,
		LoanID:       loan.LoanID,
	}, nil
}
