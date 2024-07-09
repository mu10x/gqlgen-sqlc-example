package gqlgen

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mu10x/gqlgen-sqlc-example/pg"
)

func NewHandler(repo pg.Repository) http.Handler {
	return handler.New(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repository: repo,
		},
	}))
}

func NewPlaygroundHandler(endpoint string) http.Handler {
	return playground.Handler("GraphQL Playground", endpoint)
}
