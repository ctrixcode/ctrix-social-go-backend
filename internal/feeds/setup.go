package feeds

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// SetupFeeds initializes and registers all feed-related components and routes.
func SetupFeeds(router chi.Router, db *sqlx.DB) error {
	// Initialize repository
	feedRepo := NewPGFeedRepository(db)

	// Initialize service
	feedService := NewFeedService(feedRepo)

	// Initialize handler
	feedHandler := NewFeedHandler(feedService)

	// Register routes
	RegisterFeedRoutes(router, feedHandler)

	return nil
}
