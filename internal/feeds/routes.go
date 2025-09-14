package feeds

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterFeedRoutes(router chi.Router, handler *FeedHandler) {
	router.Route("/feed", func(r chi.Router) {
		r.Get("/posts", middleware.WrapHandler(handler.GetFeedHandler))
	})
}
