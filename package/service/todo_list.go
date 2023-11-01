package service

import (
	ToDoList "ToDO"
	"ToDO/package/repository"
)

type TodoListService struct {
	repo repository.TogoList
}

func NewTodoListService(repo repository.TogoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list ToDoList.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]ToDoList.TodoList, error) {
	return s.repo.GetAll(userId)
}
func (s *TodoListService) GetById(userId, listId int) (ToDoList.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
func (s *TodoListService) Update(userId, listId int, input ToDoList.UpdateListInput) error {
	if err := input.Validate(); err != nil {

	}
	return s.repo.Update(userId, listId, input)
}
