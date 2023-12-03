package controller

import (
	"errors"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/model"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func ChangeTodoStatus(id string, title string, done bool) (*Todo, error) {
	// IDが空の場合は新規TODO作成
	if id == "" {
		todo, err := model.NewTodo(title, done)
		if err != nil {
			return nil, err
		}

		createdTodo, err := model.CreateTodo(todo)
		if err != nil {
			return nil, err
		}

		return &Todo{
			ID:    createdTodo.ID,
			Title: createdTodo.Title,
			Done:  createdTodo.Done,
		}, nil
	}

	if title != "" {
		return nil, errors.New("新規で作成するTODOではない場合は、タイトルの変更ができません。")
	}

	todo, err := model.GetTodoByID(id)
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
	updatedTodo, err := model.UpdateTodoStatus(todo)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:    updatedTodo.ID,
		Title: updatedTodo.Title,
		Done:  updatedTodo.Done,
	}, nil
}
