package posts

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/jwt"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterPostRoutes(router chi.Router, handler *PostHandler) {
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		panic(err)
	}
	router.Use(middleware.AuthMiddleware(jwtService))
	router.Post("/", middleware.WrapHandler(handler.CreatePost))
	router.Put("/{id}", middleware.WrapHandler(handler.UpdatePost))
}
