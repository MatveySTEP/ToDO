package repository

import (
	ToDoList "ToDO"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ToDoList.User) (int, error)
	GetUser(username, password string) (ToDoList.User, error)
}
type TodoList interface {
	Create(userId int, list ToDoList.TodoList) (int, error)
	GetAll(userId int) ([]ToDoList.TodoList, error)
	GetById(userId, listId int) (ToDoList.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input ToDoList.UpdateListInput) error
}
type TodoItem interface {
	Create(listId int, list ToDoList.TodoItem) (int, error)
	GetAll(userId, listId int) ([]ToDoList.TodoItem, error)
	GetById(userId, itemId int) (ToDoList.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemid int, input ToDoList.UpdateItemInput) error
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
