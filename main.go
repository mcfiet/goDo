package main

import (
	"log"
	"net/http"

	"github.com/mcfiet/goDo/db"
	router "github.com/mcfiet/goDo/todo/routing"
)

func main() {
	db.Init()
	r := router.TodoRouter()
	log.Println("Server starting on :3000")
	http.ListenAndServe(":3000", r)
}
