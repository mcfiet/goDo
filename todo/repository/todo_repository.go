package repository

import (
	"devoniq/goDo/db"
	"devoniq/goDo/todo/model"
)

func GetAllTodos() ([]model.Todo, error) {
	database := db.GetDB()
	var todos []model.Todo
	result := database.Find(&todos)

	return todos, result.Error
}

func CreateTodo(todo *model.Todo) error {
	database := db.GetDB()

	result := database.Create(todo)

	return result.Error
}
