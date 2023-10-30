package service

import (
	ToDoList "ToDO"
	"ToDO/package/repository"
)

type Authorization interface {
	CreateUser(user ToDoList.User) (int, error)
	GenerateToken(username, password string) (string, error)
}
type TogoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization
	TogoList
	TodoItem
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
