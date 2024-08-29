package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"todo/internal/handlers"
)

func main() {

	// Вынес подключение к Базе данных из файла сюда и снес файл database.go
	db, err := initDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	handler := handlers.New(db)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func initDB() (*sql.DB, error) {
	connect := "user=postgres password=12345 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Успешное подключение к БД")
	return db, nil
}
