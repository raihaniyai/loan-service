package loan

import (
	"errors"
	"time"

	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/constant"
)

const (
	maxInterestRate = 0.99 // 99%
	maxROI          = 0.99 // 99%
)

func (svc *service) CreateLoan(request CreateLoanRequest) (CreateLoanResult, error) {
	if request.User.Role != constant.UserRoleBorrower {
		err := errors.New("user is not a borrower")
		return CreateLoanResult{}, err
	}

	loan, err := svc.loanRepository.GetLoanByBorrowerIDAndNotInStatus(request.User.UserID, constant.LoanStatusClosed)
	if err != nil {
		return CreateLoanResult{}, err
	}

	// assumption: one borrower can only have one loan at a time
	if loan != nil {
		return CreateLoanResult{}, errors.New("user already has an active loan")
	}

	if request.InterestRate > maxInterestRate || request.ReturnOnInvestment > maxROI {
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

	var loanID int64
	loanID, err = svc.loanRepository.SetLoan(loan)
	if err != nil {
		return CreateLoanResult{}, err
	}

	return CreateLoanResult{
		LoanID: loanID,
	}, nil
}
