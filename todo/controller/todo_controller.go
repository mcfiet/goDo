package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mcfiet/goDo/todo/model"
	todoService "github.com/mcfiet/goDo/todo/service"
)

type TodoController struct {
	service *todoService.TodoService
}

func NewTodoController(service *todoService.TodoService) *TodoController {
	return &TodoController{service}
}

func (controller *TodoController) GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Println(r)
	w.Write([]byte("This is my api"))
}

func (controller *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	entries, err := controller.service.GetAllTodos()
	if err != nil {
		http.Error(w, "Fehler beim Holen der Daten", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		http.Error(w, "Fehler beim Codieren der Todos", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (controller *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	err := controller.service.CreateTodo(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (controller *TodoController) GetTodoById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	entry, err := controller.service.GetTodoById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(&entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (controller *TodoController) UpdateTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todo.ID = parsedUUID
	if err := controller.service.UpdateTodoById(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}
