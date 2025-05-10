package internal

import (
	"net/http"
	"task-manager/internal/handler"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Serve static files
	fileServer := http.FileServer(http.Dir("static"))
	r.Handle("/*", fileServer)

	// API routes
	r.Get("/tasks", handler.GetTasks)
	r.Post("/tasks", handler.CreateTask)
	r.Patch("/tasks/{id}", handler.UpdateTask)
	return r
}
