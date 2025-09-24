package users_data

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, handler *DataHandler) {
	router.Put("/follow/{follower_id}", middleware.WrapHandler(handler.Follow))
	router.Delete("/follow/{follower_id}", middleware.WrapHandler(handler.UnFollow))
}
