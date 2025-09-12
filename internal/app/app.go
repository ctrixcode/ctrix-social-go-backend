package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"

	// Import your modules here
	"github.com/mcctrix/ctrix-social-go-backend/internal/auth"
	"github.com/mcctrix/ctrix-social-go-backend/internal/healthcheck"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/database" // Import database package
)

// Application holds all application-wide dependencies
type Application struct {
	Router *chi.Mux
	DB     *sqlx.DB
	// Add other dependencies like Logger, Config, etc. here
}

// NewApplication creates and initializes a new Application instance
func NewApplication() (*Application, error) {
	authRepo := auth.NewAuthRepository(database.DBConnection())
	user, err := authRepo.GetUserByID("1")
	fmt.Println(user)
	fmt.Println(err)
	app := &Application{
		Router: chi.NewRouter(),
	}

	// Initialize database connection
	app.DB = database.DBConnection() // Assuming DBConnection returns *sql.DB

	return app, nil
}

// SetupRoutes configures all application routes, including versioning
func (app *Application) SetupRoutes() {
	// Global middleware
	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)

	app.Router.Route("/api", func(r chi.Router) {

		// API Version 1 routes
		app.Router.Group(func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				// Add other v1 module routes here:
				// users.RegisterRoutes(r, app.DB) // Example for users module
				// posts.RegisterRoutes(r, app.DB) // Example for posts module
			})
		})

		// API Version 2 routes (example)
		app.Router.Group(func(r chi.Router) {
			r.Route("/v2", func(r chi.Router) {
				// healthcheck.RegisterRoutesV2(r) // If healthcheck has a v2
				// users.RegisterRoutesV2(r, app.DB)
			})
		})

		// Add a root handler for unmatched routes
		app.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not Found", http.StatusNotFound)
		})

		healthcheck.RegisterRoutes(r)
	})
}

// Serve starts the HTTP server
func (app *Application) Serve() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	fmt.Printf("Server starting on port %s\n", port)
	return http.ListenAndServe(":"+port, app.Router)
}
