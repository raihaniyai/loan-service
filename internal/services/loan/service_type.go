package loan

import "loan-service/internal/entity"

type CreateLoanRequest struct {
	User               entity.User
	PrincipalAmount    int64
	InterestRate       float32
	ReturnOnInvestment float32
}

type CreateLoanResult struct {
	LoanID int64
}
