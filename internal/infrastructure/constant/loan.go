package constant

type LoanStatus int

const (
	LoanStatusProposed  LoanStatus = 10
	LoanStatusApproved  LoanStatus = 20
	LoanStatusRejected  LoanStatus = 21
	LoanStatusInvested  LoanStatus = 30
	LoanStatusDisbursed LoanStatus = 40
)

var (
	LoanStatusText = map[LoanStatus]string{
		LoanStatusProposed:  "Proposed",
		LoanStatusApproved:  "Approved",
		LoanStatusRejected:  "Rejected",
		LoanStatusInvested:  "Invested",
		LoanStatusDisbursed: "Disbursed",
	}
)
