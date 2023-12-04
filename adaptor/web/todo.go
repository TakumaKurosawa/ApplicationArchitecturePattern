package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/application"
)

type TodoHandler struct {
	todoUseCase *application.TodoUseCase
}

func NewTodoHandler(todoUseCase *application.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

type ChangeTodoStatusRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (h *TodoHandler) ChangeTodoStatus(w http.ResponseWriter, r *http.Request) {
	var req ChangeTodoStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)

		return
	}

	todo, err := h.todoUseCase.ChangeTodoStatus(r.URL.Query().Get("id"), req.Title, req.Done)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating todo: %v", err), http.StatusInternalServerError)

		return
	}

	result, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting to JSON: %v", err), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}
