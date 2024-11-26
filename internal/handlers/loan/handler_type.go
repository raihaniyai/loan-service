package loan

type CreateLoanRequest struct {
	UserID             int64   `json:"-"`
	UserRole           int     `json:"-"`
	PrincipalAmount    int64   `json:"principal_amount"`
	InterestRate       float32 `json:"interest_rate"`
	ReturnOnInvestment float32 `json:"return_on_investment"`
}

type CreateLoanResponse struct {
	LoanID int64 `json:"loan_id"`
}
