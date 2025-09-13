package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/mcctrix/ctrix-social-go-backend/internal/graphql/generated"
	graphqlMiddleware "github.com/mcctrix/ctrix-social-go-backend/internal/graphql/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/graphql/resolvers"
	"github.com/mcctrix/ctrix-social-go-backend/internal/users_setting"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/jwt"
)

// SetupGraphQL initializes and registers GraphQL routes.
func SetupGraphQL(r *chi.Mux, db *sqlx.DB) error {
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		return fmt.Errorf("failed to create JWT service: %w", err)
	}

	// Initialize UserSettingService
	userSettingRepo := users_setting.NewRepository(db)
	userSettingService := users_setting.NewService(userSettingRepo)

	// Create a new GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			UserSettingService: userSettingService,
		},
	}))

	// Add the GraphQL playground
	fmt.Println("GraphQL playground available at http://localhost:4000/")
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", graphqlMiddleware.AuthMiddleware(jwtService)(srv))

	return nil
}
