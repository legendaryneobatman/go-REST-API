package repository

import (
	"github.com/jmoiron/sqlx"
	"go-server-copy/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type List interface {
	Create(userId int, list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetById(listId string) (models.List, error)
	Update(id string, input models.List) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	List
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		List:          NewListPostgres(db),
	}
}
