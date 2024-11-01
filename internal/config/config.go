package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return &Config{
		DatabaseURL: getEnv("POSTGRES_DB_URL", "postgresql://grocery_db_owner:wWy5z0donATi@ep-cold-unit-a1zy5mp1.ap-southeast-1.aws.neon.tech/grocery_db?sslmode=require"),
		Port:        getEnv("PORT", "8030"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
