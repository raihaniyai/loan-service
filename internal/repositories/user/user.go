package user

import (
	"context"
	"loan-service/internal/entity"
	"log"

	"gorm.io/gorm"
)

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	err := r.database.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println("REPO.GUEB00 | [GetUserByEmail] Error getting user:", err)
		return nil, err
	}

	return user, nil
}

func (r *repository) GetUserByUserID(ctx context.Context, userID int64) (*entity.User, error) {
	var user *entity.User
	err := r.database.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println("REPO.GUBU00 | [GetUserByUserID] Error getting user:", err)
		return nil, err
	}

	return user, nil
}

func (r *repository) SetUser(ctx context.Context, tx *gorm.DB, user *entity.User) (int64, error) {
	var result entity.User

	db := r.database
	if tx != nil {
		db = tx
	}

	err := db.Create(user).Scan(&result).Error
	if err != nil {
		log.Println("REPO.SU00 | [SetUser] Error inserting user:", err)
		return 0, err
	}

	return result.UserID, nil
}
