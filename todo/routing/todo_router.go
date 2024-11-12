package router

import (
	"net/http"

	"github.com/mcfiet/goDo/middleware"
	"github.com/mcfiet/goDo/todo/controller"

	"github.com/go-chi/chi/v5"
)

func TodoRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AuthenticationMiddleware)
	r.Get("/", controller.GetAllTodos)
	r.Post("/", controller.CreateTodo)
	r.Get("/{id}", controller.GetTodoById)
	r.Put("/{id}", controller.UpdateTodoById)

	return r
}
