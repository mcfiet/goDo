package router

import (
	"fmt"
	"net/http"

	"github.com/mcfiet/goDo/todo/controller"

	"github.com/go-chi/chi/v5"
)

func TodoRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", controller.GetHello)
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", controller.GetAllTodos)
		r.Post("/", controller.CreateTodo)
		r.Get("/{id}", controller.GetTodoById)
	})

	// Ausgabe aller registrierten Routen
	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("Registered route: %s %s\n", method, route)
		return nil
	})

	return r
}
