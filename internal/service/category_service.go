package service

import (
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(name, description string) (*model.Category, error) {
	return s.repo.Create(name, description)
}

func (s *CategoryService) GetCategories() ([]*model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetCategory(id string) (*model.Category, error) {
	return s.repo.GetByID(id)
}
