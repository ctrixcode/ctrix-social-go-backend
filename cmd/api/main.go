package main

import (
	"log"
	"os"

	"github.com/ctrixcode/ctrix-social-go-backend/internal/app"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/graphql"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvironment()

	application, err := app.NewApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	application.SetupRoutes()
	if err := graphql.SetupGraphQL(application.Router, application.DB); err != nil {
		log.Fatalf("Failed to setup GraphQL: %v", err)
	}

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
