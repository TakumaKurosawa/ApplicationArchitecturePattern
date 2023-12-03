package service

import "github.com/TakumaKurosawa/ApplicationArchitecturePattern/model"

type Todo struct {
	ID    string
	Title string
	Done  bool
}

func CreateTodo(title string, done bool) (*Todo, error) {
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

func UpdateTodoStatus(id string, done bool) (*Todo, error) {
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
