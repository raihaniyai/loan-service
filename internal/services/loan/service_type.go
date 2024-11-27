package loan

type CreateLoanRequest struct {
	UserID             int64
	UserRole           int
	PrincipalAmount    int64
	InterestRate       float32
	ReturnOnInvestment float32
}

type CreateLoanResult struct {
	LoanID int64
}
