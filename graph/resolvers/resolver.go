package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"

	"github.com/ngocxxu/grocery-store-svelte-be/graph/generated"
	"github.com/ngocxxu/grocery-store-svelte-be/graph/model"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/service"
)

//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	UserService    *service.UserService
	ProductService *service.ProductService
}

// USER
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	return (&UserResolver{r.Resolver}).CreateUser(ctx, name, email)
}

// PRODUCT
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductInput) (*model.Product, error) {
	return (&ProductResolver{r.Resolver}).CreateProduct(ctx, &input)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return (&UserResolver{r.Resolver}).Users(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return (&UserResolver{r.Resolver}).User(ctx, id)
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return (&ProductResolver{r.Resolver}).Products(ctx)
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	return (&ProductResolver{r.Resolver}).Product(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
