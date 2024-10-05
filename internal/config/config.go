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
        DatabaseURL: getEnv("DATABASE_URL", ""),
        Port:        getEnv("PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}