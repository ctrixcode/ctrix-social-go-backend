package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/mcctrix/ctrix-social-go-backend/internal/graphql/generated"
	"github.com/mcctrix/ctrix-social-go-backend/internal/graphql/resolvers"
)

func NewGraphQLServer(r *chi.Mux) {
	// Create a new GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	// Add the GraphQL playground
	fmt.Println("GraphQL playground available at http://localhost:4000/query")
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)
}
