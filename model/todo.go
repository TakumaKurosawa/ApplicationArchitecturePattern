package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func NewTodo(title string, done bool) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	if len(title) > 256 {
		return nil, errors.New("title is too long")
	}

	return &Todo{
		ID:        uuid.New().String(),
		Title:     title,
		Done:      done,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func GetTodoByID(id string) (*Todo, error) {
	row := db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var todo Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("todo was not found")
		}

		return nil, err
	}

	return &todo, nil
}

func CreateTodo(todo *Todo) (*Todo, error) {
	if _, err := db.Exec("INSERT INTO todo (`id`, `title`, `done`) VALUES (?, ?, ?)", todo.ID, todo.Title, todo.Done); err != nil {
		return nil, err
	}

	return todo, nil
}

func UpdateTodoStatus(todo *Todo) (*Todo, error) {
	if _, err := db.Exec("UPDATE todo SET done = ? WHERE id = ?", todo.Done, todo.ID); err != nil {
		return nil, err
	}

	return todo, nil
}
