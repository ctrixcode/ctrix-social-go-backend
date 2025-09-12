package healthcheck

import "github.com/go-chi/chi/v5"

// RegisterRoutes registers the healthcheck routes to the given router.
// This router will typically be a versioned sub-router (e.g., /v1).
func RegisterRoutes(r chi.Router) {
	r.Get("/healthz", Healthz)
}
