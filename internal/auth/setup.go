package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// SetupAuth initializes and registers authentication routes.
func SetupAuth(r chi.Router, db *sqlx.DB) error {
	authRepo := NewRepository(db)
	authHandler, err := NewAuthHandler(authRepo)
	if err != nil {
		return err
	}
	RegisterRoutes(r, authHandler)
	return nil
}
