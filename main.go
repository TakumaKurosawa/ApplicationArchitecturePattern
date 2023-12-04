package main

import (
	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/adaptor/web"
)

func main() {
	executor := web.NewExecutor()
	executor.Exec()
}
