package router

import (
	"devoniq/goDo/todo/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TodoRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", controller.GetHello)
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", controller.GetAllTodos)
		r.Post("/", controller.CreateTodo)
	})
	return r
}
