package infra

import (
	"database/sql"
	"errors"
	"time"
)

type TodoInfra struct {
	db *sql.DB
}

func NewTodoInfra(db *sql.DB) *TodoInfra {
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

func (i *TodoInfra) GetTodoByID(id string) (*TodoDTO, error) {
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

	return &dto, nil
}

func (i *TodoInfra) CreateTodo(dto *TodoDTO) (*TodoDTO, error) {
	if _, err := i.db.Exec("INSERT INTO todo (`id`, `title`, `done`) VALUES (?, ?, ?)", dto.ID, dto.Title, dto.Done); err != nil {
		return nil, err
	}

	return dto, nil
}

func (i *TodoInfra) UpdateTodoStatus(dto *TodoDTO) (*TodoDTO, error) {
	if _, err := i.db.Exec("UPDATE todo SET done = ? WHERE id = ?", dto.Done, dto.ID); err != nil {
		return nil, err
	}

	return dto, nil
}
