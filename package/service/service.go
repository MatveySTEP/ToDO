package service

import "ToDO/package/repository"

type Authorization interface {
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
	return &Service{}
}
