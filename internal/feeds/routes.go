package feeds

import (
	"github.com/go-chi/chi/v5"
)

func RegisterFeedRoutes(router chi.Router, handler *FeedHandler) {
	router.Route("/feed", func(r chi.Router) {
		r.Get("/posts", handler.GetFeedHandler)
	})
}
