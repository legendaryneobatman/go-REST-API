package service

import (
	"go-server-copy/models"
	"go-server-copy/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	Create(userId int, list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetById(listId string) (models.List, error)
	Update(listId string, input models.List) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	List
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
	}
}
