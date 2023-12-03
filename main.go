package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todo/done", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Todo Done")
	})

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error:", err)
	}
}
