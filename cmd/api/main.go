package main

import (
	"log/slog"
	"os"

	"github.com/ctrixcode/ctrix-social-go-backend/internal/app"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/graphql"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	logger.InitLogger(cfg.Log)

	application, err := app.NewApplication(cfg)
	if err != nil {
		slog.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}

	application.SetupRoutes()
	if err := graphql.SetupGraphQL(application.Router, application.DB, cfg); err != nil {
		slog.Error("Failed to setup GraphQL", "error", err)
		os.Exit(1)
	}

	if err := application.Serve(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
