package service

import (
	ToDoList "ToDO"
	"ToDO/package/repository"
)

type Authorization interface {
	CreateUser(user ToDoList.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TogoList interface {
	Create(userId int, list ToDoList.TodoList) (int, error)
	GetAll(userId int) ([]ToDoList.TodoList, error)
	GetById(userId, listId int) (ToDoList.TodoList, error)
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
		TogoList:      NewTodoListService(repos.TogoList),
	}
}
