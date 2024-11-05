package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/db"
	loginRouter "github.com/mcfiet/goDo/login/routing"
	todoRouter "github.com/mcfiet/goDo/todo/routing"
)

func main() {
	db.Init()
	r := chi.NewRouter()
	todoRouter := todoRouter.TodoRouter()
	loginRouter := loginRouter.LoginRouter()

	r.Mount("/todos", todoRouter)
	r.Mount("/login", loginRouter)
	log.Println("Server starting on :3000")
	http.ListenAndServe(":3000", r)
}
