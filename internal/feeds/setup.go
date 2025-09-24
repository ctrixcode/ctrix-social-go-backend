package feeds

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Setup(router chi.Router, db *sqlx.DB) {
	// Initialize repository
	feedRepo := NewPGFeedRepository(db)

	// Initialize service
	feedService := NewFeedService(feedRepo)

	// Initialize handler
	feedHandler := NewFeedHandler(feedService)

	// Register routes
	RegisterFeedRoutes(router, feedHandler)

}
