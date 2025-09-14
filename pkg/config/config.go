package config

import "os"

// CloudinaryConfig holds the configuration for Cloudinary.
type CloudinaryConfig struct {
	CloudName string
	APIKey    string
	APISecret string
}

// LoadCloudinaryConfig loads Cloudinary configuration from environment variables.
func LoadCloudinaryConfig() CloudinaryConfig {
	return CloudinaryConfig{
		CloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		APIKey:    os.Getenv("CLOUDINARY_API_KEY"),
		APISecret: os.Getenv("CLOUDINARY_API_SECRET"),
	}
}
