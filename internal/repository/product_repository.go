package repository

import (
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	result := r.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *ProductRepository) GetAll() ([]*model.Product, error) {
	var products []*model.Product
	result := r.DB.Preload("WeightOptions.Unit").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id string) (*model.Product, error) {
	var product model.Product
	result := r.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
