package main

import (
	"log"

	"github.com/ctrixcode/ctrix-social-go-backend/internal/app"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/graphql"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	application.SetupRoutes()
	if err := graphql.SetupGraphQL(application.Router, application.DB, cfg); err != nil {
		log.Fatalf("Failed to setup GraphQL: %v", err)
	}

	if err := application.Serve(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
