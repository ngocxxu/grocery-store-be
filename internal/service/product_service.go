package service

import (
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/repository"
)

type ProductService struct {
    repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
    return s.repo.Create(product)
}

func (s *ProductService) GetProducts() ([]*model.Product, error) {
    return s.repo.GetAll()
}

func (s *ProductService) GetProduct(id string) (*model.Product, error) {
	return s.repo.GetByID(id)
}