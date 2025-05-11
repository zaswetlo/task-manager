package internal

import (
	"net/http"
	"task-manager/internal/frontend"
	"task-manager/internal/handler"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Initialize frontend handler
	frontendHandler, err := frontend.NewHandler()
	if err != nil {
		panic(err)
	}

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/tasks", handler.GetTasks)
		r.Post("/tasks", handler.CreateTask)
		r.Patch("/tasks/{id}", handler.UpdateTask)
		r.Put("/tasks/{id}", handler.UpdateTaskDetails)
		r.Delete("/tasks/{id}", handler.DeleteTask)
		r.Get("/tasks/progress", handler.GetTaskProgress)
	})

	// Serve static files
	r.Handle("/static/*", frontendHandler.ServeStatic())

	// Serve index page for all other routes
	r.Get("/*", frontendHandler.ServeIndex)

	return r
}
