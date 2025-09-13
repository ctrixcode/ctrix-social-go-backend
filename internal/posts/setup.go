package posts

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func SetupPosts(router chi.Router, db *sqlx.DB) error {
	repo := NewRepository(db)
	service := NewService(repo)
	validator := validator.New()
	handler := NewPostHandler(service, validator)

	RegisterPostRoutes(router, handler)
	return nil
}
