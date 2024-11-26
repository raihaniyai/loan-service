package action

import "loan-service/internal/entity"

type ApproveLoanRequest struct {
	User        entity.User
	LoanID      int64
	DocumentURL string
}

type ApproveLoanResult struct {
	LoanID int64
}
