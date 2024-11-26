package action

import (
	"context"
	"log"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

func (r *repository) SetAction(ctx context.Context, tx *gorm.DB, action *entity.Action) (int64, error) {
	var result entity.Action

	db := r.database
	if tx != nil {
		db = tx
	}

	err := db.Create(action).Scan(&result).Error
	if err != nil {
		log.Println("REPO.SA00 | [SetAction] Error inserting action:", err)
		return 0, err
	}

	return result.ActionID, nil
}
