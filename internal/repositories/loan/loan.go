package loan

import (
	"context"
	"log"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

func (r *repository) GetLoanByBorrowerIDAndNotInStatuses(ctx context.Context, userID int64, statuses []int) (*entity.Loan, error) {
	var loan *entity.Loan
	err := r.database.Where("borrower_id = ? AND status NOT IN (?)", userID, statuses).First(&loan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println("REPO.GLBBIDANIS00 | [GetLoanByBorrowerIDAndNotInStatus] Error getting loan:", err)
		return nil, err
	}

	return loan, nil
}

func (r *repository) GetLoanByID(ctx context.Context, loanID int64) (*entity.Loan, error) {
	var loan *entity.Loan
	err := r.database.Where("loan_id = ?", loanID).First(&loan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println("REPO.GLBI00 | [GetLoanByID] Error getting loan:", err)
		return nil, err
	}

	return loan, nil
}

func (r *repository) SetLoan(ctx context.Context, tx *gorm.DB, loan *entity.Loan) (int64, error) {
	var result entity.Loan
	err := r.database.Create(loan).Scan(&result).Error
	if err != nil {
		log.Println("REPO.SL00 | [SetLoan] Error inserting loan:", err)
		return 0, err
	}
	return result.LoanID, nil
}

func (r *repository) UpdateLoan(ctx context.Context, tx *gorm.DB, loan *entity.Loan) error {
	err := r.database.Save(loan).Error
	if err != nil {
		log.Println("REPO.UL00 | [UpdateLoan] Error updating loan:", err)
		return err
	}

	return nil
}
