package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/controller"
)

type ChangeTodoStatusRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func ChangeTodoStatus(w http.ResponseWriter, r *http.Request) {
	var req ChangeTodoStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)

		return
	}

	todo, err := controller.ChangeTodoStatus(r.URL.Query().Get("id"), req.Title, req.Done)
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
