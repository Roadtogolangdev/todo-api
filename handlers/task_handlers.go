package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"todo/internal/database"
	"todo/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err = database.DB.QueryRow(
		"INSERT INTO tasks (title, description, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)

	if err != nil {
		http.Error(w, "Could not create task", http.StatusInternalServerError)
		log.Printf("Failed to create task: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, title, description, due_date, created_at, updated_at FROM tasks")
	if err != nil {
		http.Error(w, "Could not get tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			http.Error(w, "Could not scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task
	err := database.DB.QueryRow(
		"SELECT id, title, description, due_date, created_at, updated_at FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not retrieve task", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task.UpdatedAt = time.Now()

	res, err := database.DB.Exec(
		"UPDATE tasks SET title = $1, description = $2, due_date = $3, updated_at = $4 WHERE id = $5",
		task.Title, task.Description, task.DueDate, task.UpdatedAt, id)

	if err != nil {
		http.Error(w, "Could not update task", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Could not delete task", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
