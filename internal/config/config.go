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
	godotenv.Load()

	return &Config{
		DatabaseURL: getEnv("POSTGRES_DB_URL", ""),
		Port:        getEnv("PORT", "8030"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
