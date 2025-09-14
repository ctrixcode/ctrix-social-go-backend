package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"

	"github.com/ctrixcode/ctrix-social-go-backend/internal/auth"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/feeds" // Added feeds import
	"github.com/ctrixcode/ctrix-social-go-backend/internal/healthcheck"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/post_comment_likes"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/post_comments"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/post_likes"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/posts"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_profile"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_setting"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/database"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/jwt"
	customMiddleware "github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
)

// Application holds all application-wide dependencies
type Application struct {
	Router             *chi.Mux
	DB                 *sqlx.DB
	UserSettingService users_setting.Service
	// Add other dependencies like Logger, Config, etc. here
}

// NewApplication creates and initializes a new Application instance
func NewApplication() (*Application, error) {
	app := &Application{
		Router: chi.NewRouter(),
	}

	// Initialize database connection
	app.DB = database.DBConnection()

	// Initialize UserSettingService
	userSettingRepo := users_setting.NewRepository(app.DB)
	app.UserSettingService = users_setting.NewService(userSettingRepo)

	return app, nil
}

// SetupRoutes configures all application routes, including versioning
func (app *Application) SetupRoutes() {
	// Global middleware
	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)
	app.Router.Use(customMiddleware.ErrorHandler) // Apply the custom error handler middleware

	app.Router.Route("/api", func(r chi.Router) {
		healthcheck.RegisterRoutes(r)

		r.Route("/v1", func(rV1 chi.Router) {
			// /v1/auth routes
			rV1.Route("/auth", func(rAuth chi.Router) {
				if err := auth.SetupAuth(rAuth, app.DB); err != nil {
					panic(err) // Handle error during setup
				}
			})

			if err := feeds.SetupFeeds(rV1, app.DB); err != nil { // Added feeds setup
				panic(err) // Handle error during setup
			}
			rV1.Route("/upload", func(rUpload chi.Router) {
				if err := users_profile.SetupUserProfile(rUpload, app.DB); err != nil {
					panic(err) // Handle error during setup
				}
			})
			rV1.Route("/post", func(rPosts chi.Router) {
				jwtService, err := jwt.NewJWTService()
				if err != nil {
					panic(err)
				}
				rPosts.Use(customMiddleware.AuthMiddleware(jwtService))
				if err := posts.SetupPosts(rPosts, app.DB); err != nil {
					panic(err) // Handle error during setup
				}
				rPosts.Route("/like", func(rPostLikes chi.Router) {
					if err := post_likes.SetupPostLikes(rPostLikes, app.DB); err != nil {
						panic(err)
					}
				})
				rPosts.Route("/comment", func(rPostComments chi.Router) {
					if err := post_comments.SetupPostComments(rPostComments, app.DB); err != nil {
						panic(err)
					}
					rPostComments.Route("/like", func(rPostCommentLikes chi.Router) {
						if err := post_comment_likes.SetupPostCommentLikes(rPostCommentLikes, app.DB); err != nil {
							panic(err)
						}
					})
				})
			})
		})

		// Add a root handler for unmatched routes
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not Found", http.StatusNotFound)
		})
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
