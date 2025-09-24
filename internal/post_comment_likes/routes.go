package post_comment_likes

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterPostCommentLikeRoutes(router chi.Router, handler *PostCommentLikeHandler) {
	router.Post("/{id}", middleware.WrapHandler(handler.LikeComment))
	router.Delete("/{id}", middleware.WrapHandler(handler.UnlikeComment))
}
