package api

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todos")
}
