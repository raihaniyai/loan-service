package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"loan-service/configs"
)

var DB *sql.DB

func InitializeDatabase(config configs.DatabaseConfig) error {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		"disable",
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
