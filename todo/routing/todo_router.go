package router

import (
	"net/http"

	"github.com/mcfiet/goDo/todo/controller"
	"github.com/mcfiet/goDo/todo/middleware"

	"github.com/go-chi/chi/v5"
)

func TodoRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.UseIsAuthenticated)

	r.Get("/", controller.GetAllTodos)
	r.Post("/", controller.CreateTodo)
	r.Get("/{id}", controller.GetTodoById)
	r.Put("/{id}", controller.UpdateTodoById)

	return r
}
