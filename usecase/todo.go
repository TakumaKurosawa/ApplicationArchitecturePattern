package usecase

import (
	"errors"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/domain"
)

type TodoUseCase struct {
	TodoDomain *domain.TodoDomain
}

func NewTodoUseCase(todoDomain *domain.TodoDomain) *TodoUseCase {
	return &TodoUseCase{
		TodoDomain: todoDomain,
	}
}

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (u *TodoUseCase) ChangeTodoStatus(id string, title string, done bool) (*Todo, error) {
	// IDが空の場合は新規TODO作成
	if id == "" {
		t, err := u.TodoDomain.CreateTodo(title, done)
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

	t, err := u.TodoDomain.UpdateTodoStatus(id, done)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:    t.ID,
		Title: t.Title,
		Done:  t.Done,
	}, nil
}
