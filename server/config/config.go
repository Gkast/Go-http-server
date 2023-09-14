package config

import (
	"os"
)

// Config holds application configuration settings
type Config struct {
	ServerAddress string
	StaticDir     string
	// Add more configuration options as needed
}

// NewConfig creates a new configuration with default values
func NewConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"), // Default to port 8080 if not specified
		StaticDir:     getEnv("STATIC_DIR", "static"),    // Default to "static" directory
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
