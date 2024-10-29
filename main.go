package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", getHello)
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", getTodos)
	})
	http.ListenAndServe(":3000", r)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Println(r)
	io.WriteString(w, "This is my api")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getTodos")
	io.WriteString(w, "Todos")
}
