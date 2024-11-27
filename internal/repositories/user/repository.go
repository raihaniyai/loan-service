package user

import (
	"context"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByUserID(ctx context.Context, userID int64) (*entity.User, error)
	SetUser(ctx context.Context, tx *gorm.DB, user *entity.User) (int64, error)
}

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) Repository {
	return &repository{
		database: database,
	}
}
