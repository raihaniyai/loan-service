package configs

import (
	"os"

	"loan-service/internal/infrastructure/constant"
)

type Config struct {
	Database DatabaseConfig
	NSQ      NSQConfig
}

type DatabaseConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

type NSQConfig struct {
	NSQDAddress    string
	LookupDAddress string
	Topic          string
	Channel        string
}

func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
		},
		NSQ: NSQConfig{
			NSQDAddress:    "localhost:4150",
			LookupDAddress: "localhost:4161",
			Topic:          constant.NSQTopicLoanInvestmentCompleted,
			Channel:        constant.NSQLoanChannel,
		},
	}
}
