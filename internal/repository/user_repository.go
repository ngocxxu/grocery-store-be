package repository

import (
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) Create(name, email string) (*model.User, error) {
    user := &model.User{Name: name, Email: email}
    result := r.DB.Create(user)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
    var users []*model.User
    result := r.DB.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {
    var user model.User
    result := r.DB.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}