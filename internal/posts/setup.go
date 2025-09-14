package posts

import (
	"github.com/ctrixcode/ctrix-social-go-backend/internal/auth"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/cloudinary"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func Setup(router chi.Router, db *sqlx.DB, cld *cloudinary.Service) {
	repo := NewRepository(db)
	authRepo := auth.NewRepository(db)
	service := NewService(repo, cld, authRepo)
	validator := validator.New()
	handler := NewPostHandler(service, validator)

	RegisterPostRoutes(router, handler)
}
