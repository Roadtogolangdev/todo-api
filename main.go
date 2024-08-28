package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo/handlers"
	"todo/internal/database"
)

func main() {

	database.Init()
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
