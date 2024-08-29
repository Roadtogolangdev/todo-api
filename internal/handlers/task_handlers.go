package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"todo/internal/models"
)

type Handler struct {
	db *sql.DB
}

func New(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	// Получаем данные из запроса и выводим ошибку с сообщением, добавил ее так как было тяжено отловить Invalid output и еще зачем то добавил вывод кода ошибки
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		errorMessage := fmt.Sprintf("Невозможно создать задачу: %d %s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		http.Error(w, errorMessage, http.StatusInternalServerError)
		log.Printf("Ошибка при создании задачи: %v", err)
		return
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err = h.db.QueryRow(
		"INSERT INTO tasks (title, description, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)

	w.WriteHeader(http.StatusCreated)
	// Добавил добавление заголовков согласно тех.заданию чтобы каждый раз не добавлять ручками в Postman
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Printf("Неудалось закодировать ответ: %v", err)
		http.Error(w, "Нельзя закодировать ответ", http.StatusInternalServerError)
	}
}
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT id, title, description, due_date, created_at, updated_at FROM tasks")
	if err != nil {
		http.Error(w, "Не могу получить список задач", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			http.Error(w, "Не могу прочитать задачу", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Could not encode tasks to JSON", http.StatusInternalServerError)
	}
}
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task

	err := h.db.QueryRow(
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
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Could not encode task to JSON", http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task.UpdatedAt = time.Now()

	res, err := h.db.Exec(
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
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Could not encode task to JSON", http.StatusInternalServerError)
	}
}
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := h.db.Exec("DELETE FROM tasks WHERE id = $1", id)
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
