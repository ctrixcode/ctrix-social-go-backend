package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/graphql/generated"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/graphql/resolvers"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/post_comments"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/posts"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_data"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_profile"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_setting"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/cloudinary"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/jwt"
	graphqlMiddleware "github.com/ctrixcode/ctrix-social-go-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// SetupGraphQL initializes and registers GraphQL routes.
func SetupGraphQL(r *chi.Mux, db *sqlx.DB) error {
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		return fmt.Errorf("failed to create JWT service: %w", err)
	}
	cloudinaryService, err := cloudinary.NewService(config.LoadCloudinaryConfig())
	if err != nil {
		fmt.Println(err)
		return errors.InternalServerError(errors.ErrSomethingWentWrong)
	}

	// Initialize UserSettingService
	userSettingRepo := users_setting.NewRepository(db)
	userSettingService := users_setting.NewService(userSettingRepo)

	// Initialize UserDataService
	userDataRepo := users_data.NewRepository(db)
	userDataService := users_data.NewService(userDataRepo)

	// Initialize UserProfileService
	userProfileRepo := users_profile.NewRepository(db)
	userProfileService := users_profile.NewService(userProfileRepo, cloudinaryService)

	// Initialize PostService
	postRepo := posts.NewRepository(db)
	postService := posts.NewService(postRepo)

	// Initialize PostCommentService
	postCommentRepo := post_comments.NewRepository(db)
	postCommentService := post_comments.NewPostCommentService(postCommentRepo)

	// Create a new GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			UserSettingService: userSettingService,
			UserDataService:    userDataService,
			PostService:        postService,
			UserProfileService: userProfileService,
			CommentService:     postCommentService,
		},
	}))

	// Add the GraphQL playground
	fmt.Println("GraphQL playground available at http://localhost:4000/")
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", graphqlMiddleware.AuthMiddleware(jwtService)(srv))

	return nil
}
