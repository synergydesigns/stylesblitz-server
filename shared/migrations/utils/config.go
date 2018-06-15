package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string

	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	MigrationPath string
}

// LoadConfig loads db configuration from the config.toml file
func LoadConfig() *Config {
	config := viper.New()
	config.SetConfigName("Config")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error context file: %s \n", err)
	}

	return &Config{
		AppName: config.Get("app-name").(string),

		DBHost:        config.Get("db.host").(string),
		DBPort:        config.Get("db.port").(string),
		DBUser:        config.Get("db.user").(string),
		DBPassword:    config.Get("db.password").(string),
		DBName:        config.Get("db.dbname").(string),
		MigrationPath: config.Get("migration-path").(string),
	}
}
