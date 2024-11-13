package service

import (
	"github.com/google/uuid"
	"github.com/mcfiet/goDo/todo/model"
	todoRepository "github.com/mcfiet/goDo/todo/repository"
)

type TodoService struct {
	repo todoRepository.TodoRepository
}

func NewTodoService(repo todoRepository.TodoRepository) *TodoService {
	return &TodoService{repo}
}

func (service *TodoService) GetAllTodos() ([]model.Todo, error) {
	return service.repo.GetAllTodos()
}

func (service *TodoService) CreateTodo(todo *model.Todo) error {
	return service.repo.CreateTodo(todo)
}

func (service *TodoService) GetTodoById(id uuid.UUID) (model.Todo, error) {
	return service.repo.GetTodoById(id)
}

func (service *TodoService) UpdateTodoById(todo model.Todo) error {
	return service.repo.UpdateTodoById(todo)
}
