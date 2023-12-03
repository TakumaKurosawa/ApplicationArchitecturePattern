package controller

import (
	"errors"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/service"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func ChangeTodoStatus(id string, title string, done bool) (*Todo, error) {
	// IDが空の場合は新規TODO作成
	if id == "" {
		t, err := service.CreateTodo(title, done)
		if err != nil {
			return nil, err
		}

		return &Todo{
			ID:    t.ID,
			Title: t.Title,
			Done:  t.Done,
		}, nil
	}

	if title != "" {
		return nil, errors.New("新規で作成するTODOではない場合は、タイトルの変更ができません。")
	}

	t, err := service.UpdateTodoStatus(id, done)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:    t.ID,
		Title: t.Title,
		Done:  t.Done,
	}, nil
}