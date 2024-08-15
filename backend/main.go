package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", api.handler)
	// http.HandleFunc("/todos", api.getTodos)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
