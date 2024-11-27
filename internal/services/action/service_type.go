package action

import "loan-service/internal/entity"

type UpdateLoanRequest struct {
	User        entity.User
	LoanID      int64
	DocumentURL string
	ActionType  int
}

type UpdateLoanResult struct {
	LoanID int64
}

type InvestLoanRequest struct {
	User             entity.User
	LoanID           int64
	InvestmentAmount int64
}

type InvestLoanResult struct {
	InvestmentID int64
	LoanID       int64
}
