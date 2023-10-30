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
