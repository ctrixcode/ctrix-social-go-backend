package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application-wide configurations.
type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	JWT        JWTConfig
	Cloudinary CloudinaryConfig
}

// ServerConfig holds server-specific configurations.
type ServerConfig struct {
	Port string
}

// DatabaseConfig holds database-specific configurations.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig holds JWT-specific configurations.
type JWTConfig struct {
	Secret string
}

// CloudinaryConfig holds the configuration for Cloudinary.
type CloudinaryConfig struct {
	CloudName string
	APIKey    string
	APISecret string
}

// LoadConfig loads all configurations from environment variables.
func LoadConfig() (*Config, error) {
	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("No .env file found or failed to load: %v\n", err)
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "4000"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("POSTGRES_HOST", ""),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USERNAME", ""),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			DBName:   getEnv("POSTGRES_DB", ""),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", ""), // Assuming JWT_SECRET is an environment variable
		},
		Cloudinary: CloudinaryConfig{
			CloudName: getEnv("CLOUDINARY_CLOUD_NAME", ""),
			APIKey:    getEnv("CLOUDINARY_API_KEY", ""),
			APISecret: getEnv("CLOUDINARY_API_SECRET", ""),
		},
	}

	// Validate essential configurations
	if cfg.Database.Host == "" {
		return nil, fmt.Errorf("POSTGRES_HOST environment variable not set")
	}
	if cfg.Database.User == "" {
		return nil, fmt.Errorf("POSTGRES_USERNAME environment variable not set")
	}
	if cfg.Database.Password == "" {
		return nil, fmt.Errorf("POSTGRES_PASSWORD environment variable not set")
	}
	if cfg.Database.DBName == "" {
		return nil, fmt.Errorf("POSTGRES_DB environment variable not set")
	}
	if cfg.JWT.Secret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable not set")
	}
	if cfg.Cloudinary.CloudName == "" || cfg.Cloudinary.APIKey == "" || cfg.Cloudinary.APISecret == "" {
		return nil, fmt.Errorf("Cloudinary credentials (CLOUDINARY_CLOUD_NAME, CLOUDINARY_API_KEY, CLOUDINARY_API_SECRET) not fully set")
	}

	return cfg, nil
}

// getEnv retrieves the value of an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
