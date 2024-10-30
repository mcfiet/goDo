package main

import (
	"devoniq/goDo/db"
	router "devoniq/goDo/todo/routing"
	"net/http"
)

func main() {
	db.Init()
	r := router.TodoRouter()

	http.ListenAndServe(":3000", r)
}
