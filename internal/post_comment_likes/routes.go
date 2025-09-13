package post_comment_likes

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/jwt"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterPostCommentLikeRoutes(router chi.Router, handler *PostCommentLikeHandler) {
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		panic(err)
	}
	router.Use(middleware.AuthMiddleware(jwtService))
	router.Post("/{id}", middleware.WrapHandler(handler.LikeComment))
	router.Delete("/{id}", middleware.WrapHandler(handler.UnlikeComment))
}
