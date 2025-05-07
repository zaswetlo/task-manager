package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task-manager/internal"
	"testing"
)

func TestCreateAndGetTasks(t *testing.T) {
	router := internal.SetupRouter()

	// POST /tasks
	req := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"title": "Test Task"}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}

	// GET /tasks
	req = httptest.NewRequest("GET", "/tasks", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}
}
