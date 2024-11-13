package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/app"
	authRouter "github.com/mcfiet/goDo/auth/routing"
	todoRouter "github.com/mcfiet/goDo/todo/routing"
	userRouter "github.com/mcfiet/goDo/user/routing"
)

func main() {
	r := chi.NewRouter()
	app := app.InitApp()
	todoRouter := todoRouter.TodoRouter(app.TodoController)
	authRouter := authRouter.AuthRouter(app.AuthHandler, app.UserController)
	userRouter := userRouter.UserRouter(app.UserController)

	r.Mount("/", authRouter)
	r.Mount("/todos", todoRouter)
	r.Mount("/users", userRouter)
	log.Println("Server starting on :3000")

	http.ListenAndServe(":3000", r)
}
