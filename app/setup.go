package app

import (
	"github.com/mcfiet/goDo/auth/handlers"
	"github.com/mcfiet/goDo/db"
	todoController "github.com/mcfiet/goDo/todo/controller"
	todoRepository "github.com/mcfiet/goDo/todo/repository"
	todoService "github.com/mcfiet/goDo/todo/service"
	userController "github.com/mcfiet/goDo/user/controller"
	userRepository "github.com/mcfiet/goDo/user/repository"
	userService "github.com/mcfiet/goDo/user/service"
	"gorm.io/gorm"
)

type App struct {
	DB             *gorm.DB
	TodoController *todoController.TodoController
	UserController *userController.UserController
	AuthHandler    *handlers.AuthHandler
}

func InitApp() *App {
	db := db.Init()
	todoRepository := todoRepository.NewTodoRepository(db)
	todoService := todoService.NewTodoService(todoRepository)
	todoController := todoController.NewTodoController(todoService)

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)

	authHandler := handlers.NewAuthHandler(userService)
	return &App{
		DB:             db,
		TodoController: todoController,
		UserController: userController,
		AuthHandler:    authHandler,
	}
}
