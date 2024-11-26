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
