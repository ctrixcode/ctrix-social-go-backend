package users_profile

import (
	"fmt"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/cloudinary"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func SetupUserProfile(router chi.Router, db *sqlx.DB) error {
	cloudinaryService, err := cloudinary.NewService(config.LoadCloudinaryConfig())
	if err != nil {
		fmt.Println(err)
		return errors.InternalServerError(errors.ErrSomethingWentWrong)
	}

	repo := NewRepository(db)
	service := NewService(repo, cloudinaryService)
	handler := NewHandler(service)

	RegisterRoutes(router, handler)
	return nil
}
