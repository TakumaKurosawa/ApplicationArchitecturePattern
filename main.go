package main

import (
	"log"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/external/infrastructure"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/external/presentation/httpserver"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/internal/application"
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/internal/domain"
)

func main() {
	todoInfra := infrastructure.NewTodoInfra(infrastructure.ConnectDB())
	todoDomain := domain.NewTodoDomainService(todoInfra)
	todoApp := application.NewTodoUseApplication(todoDomain)

	if err := httpserver.NewExecutor().Run(todoApp); err != nil {
		log.Fatalf("HTTP server running Error:%v", err)

		return
	}
}
