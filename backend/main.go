package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GotaSuzuki/go_react/backend/api"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func InitDB() {
	var err error
	dsn := "username:password@tcp(127.0.0.1:3306)/todoapp"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 接続を確認
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database successfully!")
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	id, err := api.CreateTodo(todo.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todo.ID = int(id)
	json.NewEncoder(w).Encode(todo)
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := api.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func main() {
	dbUser := "user"
	dbPassword := "password"
	dbDatabase := "mydb"
	dbConn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dbConn)
		if err == nil {
			if err = db.Ping(); err == nil {
				log.Println("データベースに正常に接続しました")
				break
			}
		}
		log.Printf("データベースへの接続に失敗しました。エラー: %v", err)
		log.Printf("5秒後に再試行します... (試行 %d/5)", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("5回の試行後もデータベースに接続できませんでした: %v", err)
	}
	defer db.Close()

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/todos", CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todos", GetTodosHandler).Methods("GET")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
