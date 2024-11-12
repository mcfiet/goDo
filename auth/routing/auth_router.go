package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/auth/handlers"
)

func LoginRouter(handlers *handlers.AuthHandler) http.Handler {
	r := chi.NewRouter()
	log.Println("Login Router")
	r.Post("/", handlers.Login)
	return r
}
