package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/internal/app" // New import for the app package
)

func main() {
	loadEnvironment()

	application, err := app.NewApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	application.SetupRoutes()

	if err := application.Serve(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func loadEnvironment() {
	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}
	// You might want to check for critical environment variables here
	if os.Getenv("POSTGRES_HOST") == "" {
		log.Fatal("POSTGRES_HOST environment variable not set")
	}
}
