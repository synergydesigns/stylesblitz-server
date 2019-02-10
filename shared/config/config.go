package config

import (
	"os"
)

// Config struct holds all configuration
type Config struct {
	AppName string

	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	SSLMode       string
	MigrationPath string
	RootDirectory string
}

// LoadConfig loads all configuration
func LoadConfig() *Config {
	return &Config{
		AppName:       os.Getenv("APP_NAME"),
		DBHost:        os.Getenv("PG_HOST"),
		DBPort:        os.Getenv("PG_PORT"),
		DBUser:        os.Getenv("PG_USER"),
		DBPassword:    os.Getenv("PG_PASSWORD"),
		DBName:        os.Getenv("PG_DATABASE"),
		SSLMode:       os.Getenv("PG_SSL"),
		MigrationPath: os.Getenv("MIGRATION_PATH"),
		RootDirectory: os.Getenv("ROOT_DIRECTORY"),
	}
}
