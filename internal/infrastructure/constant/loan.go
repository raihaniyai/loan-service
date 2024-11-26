package constant

const (
	LoanStatusProposed  = 10
	LoanStatusApproved  = 20
	LoanStatusRejected  = 21
	LoanStatusInvested  = 30
	LoanStatusDisbursed = 40
	LoanStatusRepayment = 50
	LoanStatusClosed    = 60
)

var (
	LoanStatusText = map[int]string{
		LoanStatusProposed:  "Proposed",
		LoanStatusApproved:  "Approved",
		LoanStatusRejected:  "Rejected",
		LoanStatusInvested:  "Invested",
		LoanStatusDisbursed: "Disbursed",
	}
)