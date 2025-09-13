package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/mcctrix/ctrix-social-go-backend/internal/auth_session_tokens"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/jwt"
)

// SetupAuth initializes and registers authentication routes.
func SetupAuth(r chi.Router, db *sqlx.DB) error {
	authRepo := NewRepository(db)
	authSessionTokenRepo := auth_session_tokens.NewRepository(db)
	authSessionTokenService := auth_session_tokens.NewService(authSessionTokenRepo)

	jwtService, err := jwt.NewJWTService()
	if err != nil {
		return err
	}

	authService := NewAuthService(authRepo, authSessionTokenService, jwtService)
	authHandler, err := NewAuthHandler(authService, jwtService)
	if err != nil {
		return err
	}
	RegisterRoutes(r, authHandler)
	return nil
}
