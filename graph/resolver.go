package graph

import (
	"context"
	"strconv"

	"github.com/ngocxxu/grocery-store-svelte-be/graph/model"
)

// ... (các phần khác giữ nguyên)

func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
    user, err := r.UserService.CreateUser(name, email)
    if err != nil {
        return nil, err
    }
    return &model.User{
        ID:    strconv.FormatUint(uint64(user.ID), 10),
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
    users, err := r.UserService.GetUsers()
    if err != nil {
        return nil, err
    }
    var result []*model.User
    for _, u := range users {
        result = append(result, &model.User{
            ID:    strconv.FormatUint(uint64(u.ID), 10),
            Name:  u.Name,
            Email: u.Email,
        })
    }
    return result, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
    user, err := r.UserService.GetUser(id)
    if err != nil {
        return nil, err
    }
    return &model.User{
        ID:    strconv.FormatUint(uint64(user.ID), 10),
        Name:  user.Name,
        Email: user.Email,
    }, nil
}