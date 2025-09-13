package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/mcctrix/ctrix-social-go-backend/internal/graphql/generated"
	graphqlMiddleware "github.com/mcctrix/ctrix-social-go-backend/internal/graphql/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/graphql/resolvers"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/jwt"
)

func NewGraphQLServer(r *chi.Mux) {
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		panic(fmt.Errorf("failed to create JWT service: %w", err))
	}

	// Create a new GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	// Add the GraphQL playground
	fmt.Println("GraphQL playground available at http://localhost:4000/")
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", graphqlMiddleware.AuthMiddleware(jwtService)(srv))
}
