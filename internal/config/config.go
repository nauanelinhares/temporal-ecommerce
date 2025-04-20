package config

import (
	"fmt"
	"os"
	"strconv"
	// Optional: Use godotenv for local development to load a .env file
	// You'll need to run: go get github.com/joho/godotenv
	// "github.com/joho/godotenv"
)

// DBConfig holds database connection parameters
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // e.g., "disable", "require", "verify-full"
}

// Config holds all application configuration
type Config struct {
	Database DBConfig
	// Add other configuration sections here (e.g., ServerPort, RedisConfig)
}

// LoadConfig loads configuration from environment variables.
// It's recommended to use a .env file for local development (and use godotenv.Load()).
func LoadConfig(path string) (*Config, error) {
	// Uncomment the line below if using godotenv for .env files
	// godotenv.Load(path + "/.env") // Load .env file if present, ignore error if not found

	dbPortStr := getEnv("DATABASE_PORT", "5432") // Default PostgreSQL port
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DATABASE_PORT: %w", err)
	}

	cfg := &Config{
		Database: DBConfig{
			Host:     getEnv("DATABASE_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DATABASE_USER", "user"),
			Password: getEnv("DATABASE_PASSWORD", "password"),
			DBName:   getEnv("DATABASE_NAME", "ecommerce_db"),
			SSLMode:  getEnv("DATABASE_SSLMODE", "disable"),
		},
		// Initialize other config sections if added
	}

	return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function for error check during config loading (example)
func getEnvRequired(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("required environment variable %s not set", key)
}
