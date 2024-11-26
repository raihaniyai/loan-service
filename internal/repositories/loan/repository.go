package loan

import (
	"gorm.io/gorm"

	"loan-service/internal/entity"
)

type Repository interface {
	GetLoanByBorrowerIDAndNotInStatus(userID int64, status int) (*entity.Loan, error)
	SetLoan(loan *entity.Loan) (int64, error)
}

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) Repository {
	return &repository{
		database: database,
	}
}
