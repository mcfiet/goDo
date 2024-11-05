package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/login/controller"
)

func LoginRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", controller.GetLogin)
	r.Post("/", controller.Login)
	return r
}
