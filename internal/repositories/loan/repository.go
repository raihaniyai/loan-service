package loan

import (
	"context"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

type Repository interface {
	GetLoanByBorrowerIDAndNotInStatuses(ctx context.Context, userID int64, statuses []int) (*entity.Loan, error)
	GetLoanByID(ctx context.Context, loanID int64) (*entity.Loan, error)
	SetLoan(ctx context.Context, tx *gorm.DB, loan *entity.Loan) (int64, error)
	UpdateLoan(ctx context.Context, tx *gorm.DB, loan *entity.Loan) error
}

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) Repository {
	return &repository{
		database: database,
	}
}
