package controller

import (
	"devoniq/goDo/todo/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Println(r)
	io.WriteString(w, "This is my api")
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	entries, err := service.GetAllTodos()

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		http.Error(w, "Fehler beim Codieren der Todos", http.StatusInternalServerError)
	}
}
