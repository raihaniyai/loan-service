package entity

import "time"

type Loan struct {
	ID                 int64     `json:"loan_id"`
	BorrowerID         int64     `json:"borrower_id"`
	PrincipalAmount    int64     `json:"principal_amount"`
	InterestRate       float32   `json:"interest_rate"`        // in percentage
	ReturnOnInvestment float32   `json:"return_on_investment"` // in percentage
	AgreementLetter    string    `json:"agreement_letter,omitempty"`
	Status             int       `json:"status"`
	CreatedBy          int64     `json:"created_by"`
	UpdatedBy          int64     `json:"updated_by"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
