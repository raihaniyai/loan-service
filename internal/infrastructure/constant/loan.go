package constant

const (
	LoanStatusProposed  = 10
	LoanStatusApproved  = 20
	LoanStatusInvested  = 30
	LoanStatusDisbursed = 40
	LoanStatusRepayment = 50
	LoanStatusClosed    = 60
	LoanStatusRejected  = 61
)

var (
	LoanStatusText = map[int]string{
		LoanStatusProposed:  "Proposed",
		LoanStatusApproved:  "Approved",
		LoanStatusInvested:  "Invested",
		LoanStatusDisbursed: "Disbursed",
		LoanStatusRepayment: "Repayment",
		LoanStatusClosed:    "Closed",
		LoanStatusRejected:  "Rejected",
	}
)
