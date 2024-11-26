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

func (s *service) InvestLoan(ctx context.Context, request InvestLoanRequest) (InvestLoanResult, error) {
	if request.User.Role != constant.UserRoleInvestor {
		log.Println("SVC.IL00 | [InvestLoan] User is not an investor")
		return InvestLoanResult{}, errors.New("user is not an investor")
	}

	loan, err := s.loanRepository.GetLoanByID(ctx, request.LoanID)
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

	// assumption: user always has enough balance to invest when hitting this endpoint

	if request.InvestmentAmount > loan.PrincipalAmount {
		log.Println("SVC.IL04 | [InvestLoan] Investment amount is too high")
		return InvestLoanResult{}, errors.New("investment amount is too high")
	}

	totalCurrentInvestmentAmount, err := s.investmentRepository.GetTotalInvestmentAmountByLoanID(ctx, request.LoanID)
	if err != nil {
		log.Println("SVC.IL05 | [InvestLoan] Error getting total investment amount:", err)
		return InvestLoanResult{}, err
	}
	totalInvestmentAmount := totalCurrentInvestmentAmount + request.InvestmentAmount
	if totalInvestmentAmount > loan.PrincipalAmount {
		log.Println("SVC.IL06 | [InvestLoan] Investment amount is too high")
		return InvestLoanResult{}, errors.New("investment amount is too high")
	}

	investment, err := s.investmentRepository.GetInvestmentByLoanIDAndInvestorID(ctx, request.LoanID, request.User.UserID)
	if err != nil {
		log.Println("SVC.IL07 | [InvestLoan] Error getting investment:", err)
		return InvestLoanResult{}, err
	}

	if investment != nil {
		log.Println("SVC.IL08 | [InvestLoan] User has already invested in this loan")
		return InvestLoanResult{}, errors.New("user has already invested in this loan")
	}

	tx := s.database.BeginTx()
	defer func() {
		if err != nil {
			err = s.database.Rollback(tx)
			if err != nil {
				log.Println("SVC.IL09 | [InvestLoan] Error rolling back transaction:", err)
			}
		}
	}()

	currentTime := time.Now()
	investment = &entity.Investment{
		LoanID:           request.LoanID,
		InvestorID:       request.User.UserID,
		InvestmentAmount: request.InvestmentAmount,
		CreatedAt:        currentTime,
	}

	investment.InvestmentID, err = s.investmentRepository.SetInvestment(ctx, tx, investment)
	if err != nil {
		log.Println("SVC.IL10 | [InvestLoan] Error inserting investment:", err)
		return InvestLoanResult{}, err
	}

	if totalInvestmentAmount == loan.PrincipalAmount {
		loan.Status = constant.LoanStatusInvested
		loan.UpdatedAt = currentTime
		loan.UpdatedBy = request.User.UserID

		err = s.loanRepository.UpdateLoan(ctx, tx, loan)
		if err != nil {
			log.Println("SVC.IL11 | [InvestLoan] Error updating loan:", err)
			return InvestLoanResult{}, err
		}

		investments, err := s.investmentRepository.GetInvestmentsByLoanID(ctx, request.LoanID)
		if err != nil {
			log.Println("SVC.IL12 | [InvestLoan] Error getting investments:", err)
			return InvestLoanResult{}, err
		}

		for _, investment := range investments {
			// TODO: publish kafka message
			// send pdf documents to all investors (publish kafka message)
			fmt.Println(investment.InvestorID)
		}
	}

	errCommit := s.database.Commit(tx)
	if errCommit != nil {
		log.Println("SVC.IL13 | [InvestLoan] Error committing transaction:", errCommit)
		return InvestLoanResult{}, errCommit
	}

	return InvestLoanResult{
		InvestmentID: investment.InvestmentID,
		LoanID:       loan.LoanID,
	}, nil
}
