package cmd

import "github.com/TakumaKurosawa/ApplicationArchitecturePattern/internal/application"

type Executor interface {
	Run(todoApp application.TodoApplication) error
}
