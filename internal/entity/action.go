package entity

import "time"

type Action struct {
	ActionID    int64     `gorm:"primaryKey" json:"action_id"`
	LoanID      int64     `gorm:"not null" json:"loan_id"`
	ActionType  int       `gorm:"not null" json:"action_type"`
	ActionBy    int       `gorm:"not null" json:"action_by"` // User Role
	DocumentURL string    `json:"document_url,omitempty"`
	CreatedBy   int64     `gorm:"not null" json:"created_by"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
}
