package repository

import (
	"github.com/mcfiet/goDo/db"
	"github.com/mcfiet/goDo/todo/model"
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

func GetTodoById(id string) (model.Todo, error) {
	database := db.GetDB()
	var todo model.Todo

	result := database.First(&todo, id)

	return todo, result.Error
}
