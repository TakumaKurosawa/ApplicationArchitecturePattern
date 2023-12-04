package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/domain"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/handler"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/infra"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/usecase"
)

func main() {
	todoInfra := infra.NewTodoInfra(infra.ConnectDB())
	todoD := domain.NewTodoDomain(todoInfra)
	todoU := usecase.NewTodoUseCase(todoD)
	todoH := handler.NewTodoHandler(todoU)

	http.HandleFunc("/todo/done", todoH.ChangeTodoStatus)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error:", err)
	}
}
