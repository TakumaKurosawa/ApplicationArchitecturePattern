package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/adaptor/datastore"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/application"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/domain"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/port/executor"
)

type webExecutor struct{}

func NewExecutor() executor.Executor {
	return &webExecutor{}
}

func (e *webExecutor) Exec() {
	todoInfra := datastore.NewTodoInfra(datastore.ConnectDB())
	todoD := domain.NewTodoDomain(todoInfra)
	todoU := application.NewTodoUseCase(todoD)
	todoH := NewTodoHandler(todoU)

	http.HandleFunc("/todo/done", todoH.ChangeTodoStatus)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error:", err)
	}
}
