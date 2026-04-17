package config

import "os"

type Config struct {
	Port        string
	DatabaseDSN string
	SQLitePath  string
	LogLevel    string
	Environment string
}

func Load() Config {
	return Config{
		Port:        getEnv("APP_PORT", "8080"),
		DatabaseDSN: getEnv("DATABASE_DSN", "postgres://todo:todo@localhost:5432/todos?sslmode=disable"),
		SQLitePath:  getEnv("SQLITE_PATH", "apps/api/data/todos.db"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Environment: getEnv("APP_ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
