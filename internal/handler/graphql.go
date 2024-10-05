package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ngocxxu/grocery-store-svelte-be/graph"
	"github.com/ngocxxu/grocery-store-svelte-be/graph/generated"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/service"
)

func NewGraphQLHandler(userService *service.UserService) http.Handler {
    return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{UserService: userService}}))
}

func NewPlaygroundHandler() http.Handler {
    return playground.Handler("GraphQL playground", "/query")
}