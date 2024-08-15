package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todo1 := models.todo1
	jsonData, err := json.Marshal(todo1)
	if err != nil {
		http.Error(w, "fial to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
