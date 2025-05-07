package handler

import (
	"encoding/json"
	"net/http"
	"task-manager/internal/storage"
)

type CreateTaskRequest struct {
	Title string `json:"title"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.GetAllTasks())
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	task := storage.AddTask(req.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
