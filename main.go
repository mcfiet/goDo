package main

import (
	"github.com/mcfiet/goDo/db"
	router "github.com/mcfiet/goDo/todo/routing"
	"net/http"
)

func main() {
	db.Init()
	r := router.TodoRouter()

	http.ListenAndServe(":3000", r)
}
