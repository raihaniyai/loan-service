package entity

import "time"

type Loan struct {
	ID                 int64     `gorm:"primaryKey" json:"loan_id"`
	BorrowerID         int64     `gorm:"not null" json:"borrower_id"`
	PrincipalAmount    int64     `gorm:"not null" json:"principal_amount"`
	InterestRate       float32   `gorm:"not null" json:"interest_rate"`        // in percentage
	ReturnOnInvestment float32   `gorm:"not null" json:"return_on_investment"` // in percentage
	AgreementLetter    string    `json:"agreement_letter,omitempty"`
	Status             int       `gorm:"default:10" json:"status"`
	UpdatedBy          int64     `json:"updated_by"`
	CreatedAt          time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
