package users_profile

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers the user profile routes
func RegisterRoutes(router chi.Router, handler *Handler) {
	router.Put("/profile_picture", middleware.WrapHandler(handler.UpdateAvatar))
}
