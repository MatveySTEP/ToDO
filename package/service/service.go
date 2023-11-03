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
type TodoList interface {
	Create(userId int, list ToDoList.TodoList) (int, error)
	GetAll(userId int) ([]ToDoList.TodoList, error)
	GetById(userId, listId int) (ToDoList.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input ToDoList.UpdateListInput) error
}
type TodoItem interface {
	Create(userId, listId int, item ToDoList.TodoItem) (int, error)
	GetAll(userId, listId int) ([]ToDoList.TodoItem, error)
	GetById(userId, itemId int) (ToDoList.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, listId int, input ToDoList.UpdateItemInput) error
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
