package service

import (
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/repository"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*model.User, error) {
    return s.repo.Create(name, email)
}

func (s *UserService) GetUsers() ([]*model.User, error) {
    return s.repo.GetAll()
}

func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}