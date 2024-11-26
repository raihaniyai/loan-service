package repositories

import (
	"gorm.io/gorm"
)

type DB interface {
	BeginTx() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}

type db struct {
	database *gorm.DB
}

func New(database *gorm.DB) DB {
	return &db{
		database: database,
	}
}

func (db *db) BeginTx() *gorm.DB {
	return db.database.Begin()
}

func (db *db) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (db *db) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}
