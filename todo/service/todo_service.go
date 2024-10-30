package service

import (
	"devoniq/goDo/todo/model"
	todoRepository "devoniq/goDo/todo/repository"
)

func GetAllTodos() ([]model.Todo, error) {
	return todoRepository.GetAllTodos()
}

func CreateTodo(todo *model.Todo) error {
	return todoRepository.CreateTodo(todo)
}
