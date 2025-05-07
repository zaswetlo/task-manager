package internal

import (
	"net/http"
	"task-manager/internal/handler"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/tasks", handler.GetTasks)
	r.Post("/tasks", handler.CreateTask)
	return r
}
