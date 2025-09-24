package auth

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers the authentication routes to the given router.
func RegisterRoutes(r chi.Router, handler *AuthHandler) {
	r.Post("/register", middleware.WrapHandler(handler.Register))
	r.Post("/login", middleware.WrapHandler(handler.Login))
	r.Post("/refresh-token", middleware.WrapHandler(handler.RefreshToken))
	r.Post("/logout", middleware.WrapHandler(handler.Logout))
}
