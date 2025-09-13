package post_comment_likes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func SetupPostCommentLikes(router chi.Router, db *sqlx.DB) error {
	repo := NewRepository(db)
	service := NewService(repo)
	validator := validator.New()
	handler := NewPostCommentLikeHandler(service, validator)

	RegisterPostCommentLikeRoutes(router, handler)
	return nil
}
