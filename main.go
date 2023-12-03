package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TakumaKurosawa/ApplicationArchitecturePattern/controller"
)

func main() {
	http.HandleFunc("/todo/done", controller.ChangeTodoStatus)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error:", err)
	}
}
