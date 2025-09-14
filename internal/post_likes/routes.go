package post_likes

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterPostLikeRoutes(router chi.Router, handler *PostLikeHandler) {
	router.Post("/{id}", middleware.WrapHandler(handler.LikePost))
	router.Delete("/{id}", middleware.WrapHandler(handler.UnlikePost))
}
