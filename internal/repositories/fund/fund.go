package fund

import (
	"context"
	"log"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

func (r *repository) GetBalanceByUserID(ctx context.Context, userID int64) (int64, error) {
	var balance int64
	err := r.database.Model(&entity.Fund{}).Where("user_id = ?", userID).Select("balance").Scan(&balance).Error
	if err != nil {
		log.Println("REPO.GBBUID00 | [GetBalanceByUserID] Error getting balance:", err)
		return 0, err
	}

	return balance, nil
}

func (r *repository) SetFund(ctx context.Context, tx *gorm.DB, fund *entity.Fund) (int64, error) {
	var result entity.Fund

	db := r.database
	if tx != nil {
		db = tx
	}

	err := db.Create(fund).Scan(&result).Error
	if err != nil {
		log.Println("REPO.SF00 | [SetFund] Error inserting fund:", err)
		return 0, err
	}

	return result.FundID, nil
}

func (r *repository) UpdateBalanceByUserID(ctx context.Context, tx *gorm.DB, userID int64, balance int64) error {
	err := r.database.Model(&entity.Fund{}).Where("user_id = ?", userID).Update("balance", balance).Error
	if err != nil {
		log.Println("REPO.SBBUID00 | [SetBalanceByUserID] Error setting balance:", err)
	}

	return nil
}
