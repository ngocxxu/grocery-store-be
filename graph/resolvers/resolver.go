package graph

import (
	"context"

	"github.com/ngocxxu/grocery-store-svelte-be/graph/generated"
	"github.com/ngocxxu/grocery-store-svelte-be/graph/model"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/service"
)

type Resolver struct {
	UserService    *service.UserService
	ProductService *service.ProductService
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// USER
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	return (&UserResolver{r.Resolver}).CreateUser(ctx, name, email)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return (&UserResolver{r.Resolver}).Users(ctx)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return (&UserResolver{r.Resolver}).User(ctx, id)
}

// PRODUCT
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductInput) (*model.Product, error) {
	return (&ProductResolver{r.Resolver}).CreateProduct(ctx, &input)
}
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return (&ProductResolver{r.Resolver}).Products(ctx)
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	return (&ProductResolver{r.Resolver}).Product(ctx, id)
}
