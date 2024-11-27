package fund

import (
	"context"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

type Repository interface {
	GetBalanceByUserID(ctx context.Context, userID int64) (int64, error)
	SetFund(ctx context.Context, tx *gorm.DB, fund *entity.Fund) (int64, error)
	UpdateBalanceByUserID(ctx context.Context, tx *gorm.DB, userID int64, balance int64) error
}

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) Repository {
	return &repository{
		database: database,
	}
}
