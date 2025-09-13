package post_comments

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterPostCommentRoutes(router chi.Router, handler *PostCommentHandler) {
	router.Post("/", middleware.WrapHandler(handler.CreateComment))
	router.Put("/{id}", middleware.WrapHandler(handler.UpdateCommentByID))
}
