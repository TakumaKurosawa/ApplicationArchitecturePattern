package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func newTodo(title string, done bool) (*todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	if len(title) > 256 {
		return nil, errors.New("title is too long")
	}

	return &todo{
		ID:        uuid.New().String(),
		Title:     title,
		Done:      done,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

type changeTodoStatusRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func ChangeTodoStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var req changeTodoStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)

		return
	}

	// IDが空の場合は新規TODO作成
	if id == "" {
		t, err := newTodo(req.Title, req.Done)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating new todo: %v", err), http.StatusBadRequest)

			return
		}

		if _, err := db.Exec("INSERT INTO todo (`id`, `title`) VALUES (?, ?)", t.ID, t.Title); err != nil {
			http.Error(w, fmt.Sprintf("Error creating new todo: %v", err), http.StatusInternalServerError)

			return
		}

		result, err := json.Marshal(t)
		if err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)

		return
	}

	if req.Title != "" {
		http.Error(w, "新規で作成するTODOではない場合は、タイトルの変更ができません。", http.StatusBadRequest)

		return
	}

	row := db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	if err := row.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error reading todo: %v", err), http.StatusInternalServerError)

		return
	}

	var t todo
	if err := row.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "todo was not found", http.StatusNotFound)

			return
		}

		http.Error(w, fmt.Sprintf("Error scanning todo: %v", err), http.StatusInternalServerError)

		return
	}

	if t.Done == req.Done {
		result, err := json.Marshal(t)
		if err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)

		return
	}

	t.Done = req.Done
	if _, err := db.Exec("UPDATE todo SET done = ? WHERE id = ?", t.Done, t.ID); err != nil {
		http.Error(w, fmt.Sprintf("Error updating todo: %v", err), http.StatusInternalServerError)

		return
	}

	result, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
