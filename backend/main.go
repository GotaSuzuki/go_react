package main

import (
	"log"
	"net/http"

	"github.com/GotaSuzuki/go_react/backend/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.Handler)
	r.HandleFunc("/todos", api.GetTodos).Methods(http.MethodGet)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
