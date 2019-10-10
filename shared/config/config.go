package config

import (
	"os"
)

// Config struct holds all configuration
type Config struct {
	AppName string

	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	SSLMode              string
	MigrationPath        string
	RootDirectory        string
	AWSRegion            string
	AwsAssesKeyID        string
	AwsSecretAccessKey   string
	AwsS3Bucket          string
	AuthenticationSecret string
	PasswordSecret       string
	GoEnv                string
}

// LoadConfig loads all configuration
func LoadConfig() *Config {
	return &Config{
		AppName:              os.Getenv("APP_NAME"),
		DBHost:               os.Getenv("PG_HOST"),
		DBPort:               os.Getenv("PG_PORT"),
		DBUser:               os.Getenv("PG_USER"),
		DBPassword:           os.Getenv("PG_PASSWORD"),
		DBName:               os.Getenv("PG_DATABASE"),
		SSLMode:              os.Getenv("PG_SSL"),
		MigrationPath:        os.Getenv("MIGRATION_PATH"),
		RootDirectory:        os.Getenv("ROOT_DIRECTORY"),
		AWSRegion:            os.Getenv("AWS_DEFAULT_REGION"),
		AwsAssesKeyID:        os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey:   os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AwsS3Bucket:          os.Getenv("AWS_S3_BUCKET"),
		// AuthenticationSecret: os.Getenv("AUTHENTICATION_SECRET"),
		AuthenticationSecret: "secret",
		PasswordSecret:       "secret",
		GoEnv:                os.Getenv("GO_ENV"),
	}
}
