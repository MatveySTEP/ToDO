package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
