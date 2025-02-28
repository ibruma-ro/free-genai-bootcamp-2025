package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	DatabaseURL string
	Port        string
}

// LoadConfig loads configuration from environment variables or defaults
func LoadConfig() *Config {
	return &Config{
		DatabaseURL: getEnvOrDefault("DATABASE_URL", "words.db"),
		Port:        getEnvOrDefault("PORT", "8080"),
	}
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
