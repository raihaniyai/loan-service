package entity

import "time"

type Fund struct {
	FundID    int64     `gorm:"primaryKey" json:"fund_id"`
	UserID    int64     `gorm:"not null" json:"user_id"`
	Balance   int64     `gorm:"not null" json:"investment_amount"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
