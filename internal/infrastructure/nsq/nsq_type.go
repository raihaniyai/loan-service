package nsq

type InvestmentCompletedMessage struct {
	LoanID     int64 `json:"loan_id"`
	InvestorID int64 `json:"investor_id"`
}
