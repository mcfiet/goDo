package repository

import (
	"github.com/google/uuid"
	"github.com/mcfiet/goDo/todo/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAllTodos() ([]model.Todo, error)
	CreateTodo(todo *model.Todo) error
	GetTodoById(id string) (model.Todo, error)
	UpdateTodoById(todo model.Todo) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (repo *todoRepository) GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	result := repo.db.Find(&todos)

	return todos, result.Error
}

func (repo *todoRepository) CreateTodo(todo *model.Todo) error {
	todo.ID = uuid.New()
	result := repo.db.Create(todo)

	return result.Error
}

func (repo *todoRepository) GetTodoById(id string) (model.Todo, error) {
	var todo model.Todo

	result := repo.db.First(&todo, id)

	return todo, result.Error
}

func (repo *todoRepository) UpdateTodoById(todo model.Todo) error {
	result := repo.db.Model(&todo).Updates(todo)

	return result.Error
}
