package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ngocxxu/grocery-store-svelte-be/graph/generated"
	graph "github.com/ngocxxu/grocery-store-svelte-be/graph/resolvers"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/service"
)

func NewGraphQLHandler(userService *service.UserService, productService *service.ProductService) http.Handler {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{UserService: userService, ProductService: productService}}))
}
func NewPlaygroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/app2/query")
}
