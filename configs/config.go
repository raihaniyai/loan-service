package configs

import (
	"os"

	"loan-service/internal/infrastructure/constant"
)

type Config struct {
	Database DatabaseConfig
	NSQ      NSQConfig
	SMTP     SMTPConfig
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

type SMTPConfig struct {
	Host     string
	Port     string
	Email    string
	Password string
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
			NSQDAddress:    os.Getenv("NSQD_ADDRESS"),
			LookupDAddress: os.Getenv("NSQ_LOOKUPD_ADDRESS"),
			Topic:          constant.NSQTopicLoanInvestmentCompleted,
			Channel:        constant.NSQLoanChannel,
		},
		SMTP: SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     "587",
			Email:    os.Getenv("SMTP_EMAIL"),
			Password: os.Getenv("SMTP_PASSWORD"),
		},
	}
}
