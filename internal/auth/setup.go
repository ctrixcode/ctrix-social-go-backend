package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/mcctrix/ctrix-social-go-backend/internal/auth_session_tokens"
)

// SetupAuth initializes and registers authentication routes.
func SetupAuth(r chi.Router, db *sqlx.DB) error {
	authRepo := NewRepository(db)
	authSessionTokenRepo := auth_session_tokens.NewRepository(db)
	authHandler, err := NewAuthHandler(authRepo, authSessionTokenRepo)
	if err != nil {
		return err
	}
	RegisterRoutes(r, authHandler)
	return nil
}
