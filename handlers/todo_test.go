package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/config"
	"todo-api/models"
	"todo-api/storage"

	"github.com/gorilla/mux"
)

func TestTodoHandlers(t *testing.T) {
	config.ConnectDB()

	todo := models.Todo{Title: "Test Todo", Description: "This is a test todo", Completed: false}
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()
	CreateTodo(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", w.Code)
	}

	var createdTodo models.Todo
	json.NewDecoder(w.Body).Decode(&createdTodo)
	if createdTodo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got '%s'", createdTodo.Title)
	}

	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	w = httptest.NewRecorder()
	GetTodos(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var todos []models.Todo
	json.NewDecoder(w.Body).Decode(&todos)
	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}

	updated := models.Todo{Title: "Updated Todo", Description: "Updated description", Completed: true}
	body, _ = json.Marshal(updated)
	req = httptest.NewRequest(http.MethodPut, "/todos/"+fmt.Sprintf("%d", createdTodo.ID), bytes.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprintf("%d", createdTodo.ID)})
	w = httptest.NewRecorder()
	UpdateTodo(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var updatedTodo models.Todo
	json.NewDecoder(w.Body).Decode(&updatedTodo)
	if updatedTodo.Title != "Updated Todo" || !updatedTodo.Completed {
		t.Errorf("Update failed, got %+v", updatedTodo)
	}

	req = httptest.NewRequest(http.MethodDelete, "/todos/"+fmt.Sprintf("%d", createdTodo.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprintf("%d", createdTodo.ID)})
	w = httptest.NewRecorder()
	DeleteTodo(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("Expected status 204, got %d", w.Code)
	}

	todos, _ = storage.GetAllTodos()
	if len(todos) != 0 {
		t.Errorf("Expected 0 todos after deletion, got %d", len(todos))
	}
}
