package storage

import (
	"testing"
	"todo-api/config"
	"todo-api/models"
)

func TestCreateAndGetTodos(t *testing.T) {
	config.ConnectDB()

	todo := &models.Todo{Title: "Test Todo", Description: "This is a test todo"}

	if err := CreateTodo(todo); err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	todos, err := GetAllTodos()
	if err != nil {
		t.Fatal(err)
	}

	if len(todos) != 1 || todos[0].Title != "Test Todo" {
		t.Errorf("Expected 1 todo with title 'Test Todo', got %v", todos)
	}
}

func TestUpdateTodo(t *testing.T) {
	config.ConnectDB()

	todo := &models.Todo{Title: "Initial Title", Description: "Initial Description"}
	CreateTodo(todo)

	updatedData := &models.Todo{Title: "Updated Title", Description: "Updated Description", Completed: true}
	updatedTodo, err := UpdateTodo(todo.ID, updatedData)
	if err != nil {
		t.Fatalf("Failed to update todo: %v", err)
	}

	if updatedTodo.Title != "Updated Title" || !updatedTodo.Completed {
		t.Errorf("Todo not updated correctly, got %v", updatedTodo)
	}
}

func TestDeleteTodo(t *testing.T) {
	config.ConnectDB()

	todo := &models.Todo{Title: "To be deleted", Description: "This todo will be deleted"}
	CreateTodo(todo)

	if err := DeleteTodo(todo.ID); err != nil {
		t.Fatalf("Failed to delete todo: %v", err)
	}

	_, err := GetTodoByID(todo.ID)
	if err == nil {
		t.Errorf("Expected error when fetching deleted todo, got none")
	}
}
