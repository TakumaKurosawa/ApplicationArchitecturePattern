package infrastructure

import (
	"database/sql"
	"errors"
	"time"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/internal/domain"
)

type TodoInfra struct {
	db *sql.DB
}

func NewTodoInfra(db *sql.DB) domain.TodoRepo {
	return &TodoInfra{
		db: db,
	}
}

type TodoDTO struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (i *TodoInfra) GetTodoByID(id string) (*domain.Todo, error) {
	row := i.db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var dto TodoDTO
	if err := row.Scan(&dto.ID, &dto.Title, &dto.Done, &dto.CreatedAt, &dto.UpdatedAt, &dto.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("todo was not found")
		}

		return nil, err
	}

	return &domain.Todo{
		ID:    dto.ID,
		Title: dto.Title,
		Done:  dto.Done,
	}, nil
}

func (i *TodoInfra) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	if _, err := i.db.Exec("INSERT INTO todo (`id`, `title`, `done`) VALUES (?, ?, ?)", todo.ID, todo.Title, todo.Done); err != nil {
		return nil, err
	}

	return todo, nil
}

func (i *TodoInfra) UpdateTodoStatus(todo *domain.Todo) (*domain.Todo, error) {
	if _, err := i.db.Exec("UPDATE todo SET done = ? WHERE id = ?", todo.Done, todo.ID); err != nil {
		return nil, err
	}

	return todo, nil
}
