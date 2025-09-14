package post_likes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func Setup(router chi.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	validator := validator.New()
	handler := NewPostLikeHandler(service, validator)

	RegisterPostLikeRoutes(router, handler)

}
