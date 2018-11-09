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
	MigrationPath string
	RootDirectory string
}

// LoadConfig loads all configuration
func LoadConfig() *Config {
	// log.Println(os.Getenv("DATABASE_HOST"), "=========")

	return &Config{
		AppName:       os.Getenv("APP_NAME"),
		DBHost:        os.Getenv("DATABASE_HOST"),
		DBPort:        os.Getenv("MYSQL_PORT"),
		DBUser:        os.Getenv("MYSQL_USER"),
		DBPassword:    os.Getenv("MYSQL_PASSWORD"),
		DBName:        os.Getenv("MYSQL_DATABASE"),
		MigrationPath: os.Getenv("MIGRATION_PATH"),
		RootDirectory: os.Getenv("ROOT_DIRECTORY"),
	}
}
