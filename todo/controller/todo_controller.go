package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/todo/model"
	"github.com/mcfiet/goDo/todo/service"
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

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)

	entry, err := service.GetTodoById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(&entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}
