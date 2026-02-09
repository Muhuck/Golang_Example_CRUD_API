package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		// Log but don't fatal, as env vars might be set in the system
		println("Warning: Error loading .env file:", err.Error())
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
