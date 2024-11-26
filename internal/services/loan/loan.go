package loan

import (
	"context"
	"errors"
	"log"
	"time"

	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/constant"
)

const (
	maxInterestRate = 0.99 // 99%
	maxROI          = 0.99 // 99%
)

func (svc *service) CreateLoan(ctx context.Context, request CreateLoanRequest) (CreateLoanResult, error) {
	if request.User.Role != constant.UserRoleBorrower {
		err := errors.New("user is not a borrower")
		return CreateLoanResult{}, err
	}

	loan, err := svc.loanRepository.GetLoanByBorrowerIDAndNotInStatuses(ctx, request.User.UserID, []int{constant.LoanStatusClosed, constant.LoanStatusRejected})
	if err != nil {
		log.Println("SVC.CL00 | [CreateLoan] Error getting loan by borrower id:", err)
		return CreateLoanResult{}, err
	}

	// assumption: one borrower can only have one loan at a time
	if loan != nil {
		log.Println("SVC.CL01 | [CreateLoan] User already has an active loan")
		return CreateLoanResult{}, errors.New("user already has an active loan")
	}

	if request.InterestRate > maxInterestRate || request.ReturnOnInvestment > maxROI {
		log.Println("SVC.CL02 | [CreateLoan] Interest rate or return on investment is too high")
		return CreateLoanResult{}, errors.New("interest rate or return on investment is too high")
	}

	currentTime := time.Now()
	status := constant.LoanStatusProposed
	userID := request.User.UserID

	loan = &entity.Loan{
		BorrowerID:         userID,
		PrincipalAmount:    request.PrincipalAmount,
		InterestRate:       request.InterestRate,
		ReturnOnInvestment: request.ReturnOnInvestment,
		Status:             status,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
		UpdatedBy:          userID,
	}

	// begin transaction
	tx := svc.database.BeginTx()
	defer func() {
		if err != nil {
			err = svc.database.Rollback(tx)
			if err != nil {
				log.Println("SVC.CL03 | [CreateLoan] Error rolling back transaction:", err)
			}
		}
	}()

	var loanID int64
	loanID, err = svc.loanRepository.SetLoan(ctx, nil, loan)
	if err != nil {
		log.Println("SVC.CL04 | [CreateLoan] Error inserting loan:", err)
		return CreateLoanResult{}, err
	}

	actionCreate := &entity.Action{
		LoanID:      loanID,
		ActionType:  constant.ActionTypeCreateLoan,
		ActionBy:    request.User.Role,
		DocumentURL: "",
		CreatedAt:   currentTime,
		CreatedBy:   request.User.UserID,
	}
	_, err = svc.actionRepository.SetAction(ctx, nil, actionCreate)
	if err != nil {
		log.Println("SVC.CL05 | [CreateLoan] Error inserting action:", err)
		return CreateLoanResult{}, err
	}

	err = svc.database.Commit(tx)
	if err != nil {
		log.Println("SVC.CL06 | [CreateLoan] Error committing transaction:", err)
		return CreateLoanResult{}, err
	}

	return CreateLoanResult{
		LoanID: loanID,
	}, nil
}
