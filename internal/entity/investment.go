package entity

import "time"

type Investment struct {
	InvestmentID     int64     `gorm:"primaryKey" json:"investment_id"`
	LoanID           int64     `gorm:"not null" json:"loan_id"`
	InvestorID       int64     `gorm:"not null" json:"investor_id"`
	InvestmentAmount int64     `gorm:"not null" json:"investment_amount"`
	CreatedAt        time.Time `gorm:"not null" json:"created_at"`
}
