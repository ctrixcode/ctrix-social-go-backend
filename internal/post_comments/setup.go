package post_comments

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func SetupPostComments(router chi.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	service := NewPostCommentService(repo)
	handler := NewPostCommentHandler(service)

	RegisterPostCommentRoutes(router, handler)
}
