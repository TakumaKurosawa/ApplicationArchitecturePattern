package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TodoDomainService interface {
	CreateTodo(title string, done bool) (*Todo, error)
	UpdateTodoStatus(id string, done bool) (*Todo, error)
}

type todoDomainService struct {
	todoRepo TodoRepo
}

func NewTodoDomainService(todoRepo TodoRepo) TodoDomainService {
	return &todoDomainService{
		todoRepo: todoRepo,
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

func (d *todoDomainService) CreateTodo(title string, done bool) (*Todo, error) {
	todo, err := newTodo(title, done)
	if err != nil {
		return nil, err
	}

	createdTodo, err := d.todoRepo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:    createdTodo.ID,
		Title: createdTodo.Title,
		Done:  createdTodo.Done,
	}, nil
}

func (d *todoDomainService) UpdateTodoStatus(id string, done bool) (*Todo, error) {
	todo, err := d.todoRepo.GetTodoByID(id)
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
	updatedTodo, err := d.todoRepo.UpdateTodoStatus(todo)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:    updatedTodo.ID,
		Title: updatedTodo.Title,
		Done:  updatedTodo.Done,
	}, nil
}
