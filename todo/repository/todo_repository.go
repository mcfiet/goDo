package repository

import (
	"devoniq/goDo/db"
	"devoniq/goDo/todo/model"
	"log"
)

func GetAllTodos() ([]model.Todo, error) {
	database := db.GetDB()
	var todos []model.Todo
	result := database.Find(&todos)

	if result.Error != nil {
		log.Println("Fehler in GetAllTodos:", result.Error) // Logge den Fehler
		return nil, result.Error
	}
	return todos, result.Error
}
