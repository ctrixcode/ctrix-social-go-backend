package helpers

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// GetRequestedFields extracts the names of the fields requested in a GraphQL query.
func GetRequestedFields(ctx context.Context) []string {
	fields := graphql.CollectFieldsCtx(ctx, nil)
	fieldNames := make([]string, 0, len(fields))
	for _, f := range fields {
		fieldNames = append(fieldNames, f.Name)
	}
	return fieldNames
}
