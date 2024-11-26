package investment

import (
	"context"
	"log"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

func (r *repository) GetInvestmentByLoanIDAndInvestorID(ctx context.Context, loanID int64, investorID int64) (*entity.Investment, error) {
	var investment *entity.Investment
	err := r.database.Where("loan_id = ? AND investor_id = ?", loanID, investorID).First(&investment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println("REPO.GIBLIDAID00 | [GetInvestmentByLoanIDAndInvestorID] Error getting investment:", err)
		return nil, err
	}

	return investment, nil
}

func (r *repository) GetInvestmentsByLoanID(ctx context.Context, loanID int64) ([]entity.Investment, error) {
	var investments []entity.Investment
	err := r.database.Where("loan_id = ?", loanID).Find(&investments).Error
	if err != nil {
		log.Println("REPO.GIBLID00 | [GetInvestmentsByLoanID] Error getting investments:", err)
		return nil, err
	}

	return investments, nil
}

func (r *repository) GetTotalInvestmentAmountByLoanID(ctx context.Context, loanID int64) (int64, error) {
	var total int64
	err := r.database.
		Model(&entity.Investment{}).
		Select("COALESCE(SUM(investment_amount), 0)").
		Where("loan_id = ?", loanID).
		Scan(&total).Error
	if err != nil {
		log.Println("REPO.GTIBLIDA00 | [GetTotalInvestmentAmountByLoanID] Error getting total investment amount:", err)
		return 0, err
	}

	return total, nil
}

func (r *repository) SetInvestment(ctx context.Context, tx *gorm.DB, investment *entity.Investment) (int64, error) {
	var result entity.Investment
	err := r.database.Create(investment).Scan(&result).Error
	if err != nil {
		log.Println("REPO.SI00 | [SetInvestment] Error inserting investment:", err)
		return 0, err
	}
	return result.InvestmentID, nil
}
