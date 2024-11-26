package configs

import "os"

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
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
	}
}
