package service

import (
	todo "go-server-copy/models"
	"go-server-copy/pkg/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Create(userId int, list todo.List) (int, error) {
	return s.repo.Create(userId, list)
}
func (s *ListService) GetAll(userId int) ([]todo.List, error) {
	return s.repo.GetAll(userId)
}

func (s *ListService) GetById(listId string) (todo.List, error) {
	return s.repo.GetById(listId)
}

func (s *ListService) Update(listId string, input todo.List) error {
	return s.repo.Update(listId, input)

}
