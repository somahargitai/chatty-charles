package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	OpenAIAPIKey string
	Model        string
	Temperature  float64
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// It's okay if .env doesn't exist, we'll use environment variables
		fmt.Println("Note: .env file not found, using environment variables")
	}

	config := &Config{
		OpenAIAPIKey: os.Getenv("OPENAI_API_KEY"),
		Model:        getEnvOrDefault("MODEL", "gpt-3.5-turbo"),
		Temperature:  0.7, // Default temperature
	}

	// Validate required configuration
	if config.OpenAIAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}

	return config, nil
}

// getEnvOrDefault returns the value of the environment variable or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
