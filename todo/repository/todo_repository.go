package repository

import (
	"log"

	"github.com/google/uuid"
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
	todo.ID = uuid.New()
	log.Println("CreateTodo")
	log.Println("Todo:", todo)
	result := database.Create(todo)

	return result.Error
}

func GetTodoById(id string) (model.Todo, error) {
	database := db.GetDB()
	var todo model.Todo

	result := database.First(&todo, id)

	return todo, result.Error
}

func UpdateTodoById(todo model.Todo) error {
	database := db.GetDB()

	result := database.Model(&todo).Updates(todo)

	return result.Error
}
