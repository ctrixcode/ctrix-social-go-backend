package posts

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterPostRoutes(router chi.Router, handler *PostHandler) {
	router.Post("/", middleware.WrapHandler(handler.CreatePost))
	router.Put("/{id}", middleware.WrapHandler(handler.UpdatePost))
}
