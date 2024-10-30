package controller

import (
	"devoniq/goDo/todo/model"
	"devoniq/goDo/todo/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Println(r)
	w.Write([]byte("This is my api"))
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	entries, err := service.GetAllTodos()

	if err != nil {
		http.Error(w, "Fehler beim Holen der Daten", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		http.Error(w, "Fehler beim Codieren der Todos", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	err := service.CreateTodo(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}
