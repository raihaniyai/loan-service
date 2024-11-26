package loan

import (
	"gorm.io/gorm"

	"loan-service/internal/entity"
	"log"
)

func (r *repository) GetLoanByBorrowerIDAndNotInStatus(userID int64, status int) (*entity.Loan, error) {
	var loan *entity.Loan
	err := r.database.Where("borrower_id = ? AND status != ?", userID, status).First(&loan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println("REPO.GLBBIDANIS01 | [GetLoanByBorrowerIDAndNotInStatus] Error getting loan:", err)
		return nil, err
	}

	return loan, nil
}

func (r *repository) SetLoan(loan *entity.Loan) (int64, error) {
	var result entity.Loan
	err := r.database.Create(loan).Scan(&result).Error
	if err != nil {
		log.Println("REPO.SL01 | [SetLoan] Error inserting loan:", err)
		return 0, err
	}
	return result.ID, nil
}
