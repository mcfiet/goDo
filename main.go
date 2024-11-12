package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/app"
	loginRouter "github.com/mcfiet/goDo/auth/routing"
	todoRouter "github.com/mcfiet/goDo/todo/routing"
)

func main() {
	r := chi.NewRouter()
	app := app.InitApp()
	todoRouter := todoRouter.TodoRouter(app.TodoController)
	loginRouter := loginRouter.LoginRouter(app.AuthHandler)

	r.Mount("/login", loginRouter)
	r.Mount("/todos", todoRouter)
	log.Println("Server starting on :3000")

	http.ListenAndServe(":3000", r)
}
