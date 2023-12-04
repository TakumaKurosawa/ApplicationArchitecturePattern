package httpserver

import (
	"fmt"
	"net/http"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/internal/application"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/internal/cmd"
)

type executor struct{}

func NewExecutor() cmd.Executor {
	return &executor{}
}

func (e *executor) Run(todoApp application.TodoApplication) error {
	todoH := newTodoHandler(todoApp)

	http.HandleFunc("/todo/done", todoH.ChangeTodoStatus)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
