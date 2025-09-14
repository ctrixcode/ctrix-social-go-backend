package users_data

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Setup(router chi.Router, db *sqlx.DB, cfg *config.Config) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewDataHandler(service)

	RegisterRoutes(router, handler)
}
