package action

import (
	"context"

	"gorm.io/gorm"

	"loan-service/internal/entity"
)

type Repository interface {
	SetAction(ctx context.Context, tx *gorm.DB, action *entity.Action) (int64, error)
}

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) Repository {
	return &repository{
		database: database,
	}
}
