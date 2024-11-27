package action

type UpdateLoanRequest struct {
	UserID      int64
	UserRole    int
	LoanID      int64
	DocumentURL string
	ActionType  int
}

type UpdateLoanResult struct {
	LoanID int64
}

type InvestLoanRequest struct {
	UserID           int64
	UserRole         int
	LoanID           int64
	InvestmentAmount int64
}

type InvestLoanResult struct {
	InvestmentID int64
	LoanID       int64
}

type SendAgreementLetterRequest struct {
	LoanID     int64
	InvestorID int64
}
