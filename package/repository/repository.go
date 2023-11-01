package repository

import (
	ToDoList "ToDO"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ToDoList.User) (int, error)
	GetUser(username, password string) (ToDoList.User, error)
}
type TogoList interface {
	Create(userId int, list ToDoList.TodoList) (int, error)
	GetAll(userId int) ([]ToDoList.TodoList, error)
	GetById(userId, listId int) (ToDoList.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input ToDoList.UpdateListInput) error
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TogoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
