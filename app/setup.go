package app

import (
	"github.com/mcfiet/goDo/auth/handlers"
	"github.com/mcfiet/goDo/db"
	todoController "github.com/mcfiet/goDo/todo/controller"
	todoRepository "github.com/mcfiet/goDo/todo/repository"
	todoService "github.com/mcfiet/goDo/todo/service"
	"gorm.io/gorm"
)

type App struct {
	DB             *gorm.DB
	TodoController *todoController.TodoController
	AuthHandler    *handlers.AuthHandler
}

func InitApp() *App {
	db := db.Init()
	todoRepository := todoRepository.NewTodoRepository(db)
	todoService := todoService.NewTodoService(todoRepository)
	todoController := todoController.NewTodoController(todoService)

	authHandler := &handlers.AuthHandler{}

	return &App{
		DB:             db,
		TodoController: todoController,
		AuthHandler:    authHandler,
	}
}
