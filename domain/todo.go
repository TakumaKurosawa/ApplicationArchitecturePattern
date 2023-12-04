package domain

import (
	"errors"

	"github.com/google/uuid"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/infra"
)

type TodoDomain struct {
	todoInfra *infra.TodoInfra
}

func NewTodoDomain(todoInfra *infra.TodoInfra) *TodoDomain {
	return &TodoDomain{
		todoInfra: todoInfra,
	}
}

type Todo struct {
	ID    string
	Title string
	Done  bool
}

func newTodo(title string, done bool) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	if len(title) > 256 {
		return nil, errors.New("title is too long")
	}

	return &Todo{
		ID:    uuid.New().String(),
		Title: title,
		Done:  done,
	}, nil
}

func (d *TodoDomain) CreateTodo(title string, done bool) (*Todo, error) {
	todo, err := newTodo(title, done)
	if err != nil {
		return nil, err
	}

	createdTodo, err := d.todoInfra.CreateTodo(&infra.TodoDTO{
		ID:    todo.ID,
		Title: todo.Title,
		Done:  todo.Done,
	})
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:    createdTodo.ID,
		Title: createdTodo.Title,
		Done:  createdTodo.Done,
	}, nil
}

func (d *TodoDomain) UpdateTodoStatus(id string, done bool) (*Todo, error) {
	todo, err := d.todoInfra.GetTodoByID(id)
	if err != nil {
		return nil, err
	}

	if todo.Done == done {
		return &Todo{
			ID:    todo.ID,
			Title: todo.Title,
			Done:  todo.Done,
		}, nil
	}

	todo.Done = done
	updatedTodo, err := d.todoInfra.UpdateTodoStatus(todo)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:    updatedTodo.ID,
		Title: updatedTodo.Title,
		Done:  updatedTodo.Done,
	}, nil
}
