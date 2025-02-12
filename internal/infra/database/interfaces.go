package database

import "github.com/pansani/go-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type PostInterface interface {
	Create(post *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(post *entity.Product) error
	Delete(id string) error
}
