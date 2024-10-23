package repository

import (
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(name, description string) (*model.Category, error) {
	category := &model.Category{Name: name, Description: description}
	result := r.DB.Create(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (r *CategoryRepository) GetAll() ([]*model.Category, error) {
	var categories []*model.Category
	result := r.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id string) (*model.Category, error) {
	var category model.Category
	result := r.DB.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
