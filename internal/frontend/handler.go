package frontend

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var static embed.FS

type Handler struct {
	templates *template.Template
}

func NewHandler() (*Handler, error) {
	tmpl, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Handler{
		templates: tmpl,
	}, nil
}

func (h *Handler) ServeIndex(w http.ResponseWriter, r *http.Request) {
	if err := h.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) ServeStatic() http.Handler {
	// Create a sub-filesystem for the static directory
	fsys, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	// Create a file server that will serve files from the embedded filesystem
	fileServer := http.FileServer(http.FS(fsys))

	// Return a handler that strips the /static prefix before serving files
	return http.StripPrefix("/static/", fileServer)
}
