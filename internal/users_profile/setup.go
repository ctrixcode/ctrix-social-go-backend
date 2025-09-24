package users_profile

import (
	"log/slog"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/cloudinary"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func SetupUserProfile(router chi.Router, db *sqlx.DB, cfg *config.Config) error {
	cloudinaryService, err := cloudinary.NewService(cfg.Cloudinary)
	if err != nil {
		slog.Error("Failed to initialize cloudinary service", "error", err)
		return errors.InternalServerError(errors.ErrSomethingWentWrong)
	}

	repo := NewRepository(db)
	service := NewService(repo, cloudinaryService)
	handler := NewHandler(service)

	RegisterRoutes(router, handler)
	return nil
}
