package service

import (
	"github.com/mcfiet/goDo/todo/model"
	todoRepository "github.com/mcfiet/goDo/todo/repository"
)

func GetAllTodos() ([]model.Todo, error) {
	return todoRepository.GetAllTodos()
}

func CreateTodo(todo *model.Todo) error {
	return todoRepository.CreateTodo(todo)
}

func GetTodoById(id string) (model.Todo, error) {
	return todoRepository.GetTodoById(id)
}
