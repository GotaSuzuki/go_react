package api

import (
	"database/sql"
)

var db *sql.DB

func GetTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, title FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func CreateTodo(title string) (int64, error) {
	result, err := db.Exec("INSERT INTO todos (title) VALUES (?)", title)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
