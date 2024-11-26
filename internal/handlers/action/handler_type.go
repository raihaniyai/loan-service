package action

import "loan-service/internal/entity"

type ApproveLoanRequest struct {
	User        entity.User `json:"-"`
	LoanID      int64       `json:"-"`
	DocumentURL string      `json:"document_url"`
}

type ApproveLoanResponse struct {
	LoanID int64 `json:"loan_id"`
}

type InvestLoanRequest struct {
	User             entity.User `json:"-"`
	LoanID           int64       `json:"-"`
	InvestmentAmount int64       `json:"investment_amount"`
}

type InvestLoanResponse struct {
	InvestmentID int64 `json:"investment_id"`
	LoanID       int64 `json:"loan_id"`
}
