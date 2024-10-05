package config

import "os"

type Config struct {
    DatabaseURL string
    Port        string
}

func New() *Config {
    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=graphql_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"),
        Port:        getEnv("PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}