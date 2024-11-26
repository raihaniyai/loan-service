package investment

import (
	"context"
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetInvestmentByLoanIDAndInvestorID(ctx context.Context, loanID int64, investorID int64) (*entity.Investment, error)
	GetInvestmentsByLoanID(ctx context.Context, loanID int64) ([]entity.Investment, error)
	GetTotalInvestmentAmountByLoanID(ctx context.Context, loanID int64) (int64, error)
	SetInvestment(ctx context.Context, tx *gorm.DB, investment *entity.Investment) (int64, error)
}

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) Repository {
	return &repository{
		database: database,
	}
}
