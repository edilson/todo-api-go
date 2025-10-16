package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"todo-api/config"
	"todo-api/handlers"
	"todo-api/models"
	"todo-api/requests"
	"todo-api/storage"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func generateToken(jwtKey []byte, userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func TestTodoHandlers(t *testing.T) {
	config.ConnectDB()
	jwtKey := []byte("test_secret")

	user := models.User{Username: "bob", Password: "secret"}
	config.DB.Create(&user)

	token := generateToken(jwtKey, user.ID)

	todo := models.Todo{Title: "Test Todo", Description: "This is a test todo", Completed: false}
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	handlers.CreateTodo(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", w.Code)
	}

	var createdTodo models.Todo
	json.NewDecoder(w.Body).Decode(&createdTodo)
	if createdTodo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got '%s'", createdTodo.Title)
	}

	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	handlers.GetTodos(w, req)

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
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	handlers.UpdateTodo(w, req)

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
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	handlers.DeleteTodo(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("Expected status 204, got %d", w.Code)
	}

	todos, _ = storage.GetAllTodos()
	if len(todos) != 0 {
		t.Errorf("Expected 0 todos after deletion, got %d", len(todos))
	}
}

func TestCreateTodo(t *testing.T) {
	os.Setenv("OPENAI_API_KEY", "test_key")

	// Backup original functions so we can restore them later
	origAskMusicGenre := requests.AskMusicGenre
	origRetrievePlaylist := requests.RetrievePlaylist
	origCreateTodo := storage.CreateTodo

	defer func() {
		requests.AskMusicGenre = origAskMusicGenre
		requests.RetrievePlaylist = origRetrievePlaylist
		storage.CreateTodo = origCreateTodo
	}()

	// --- MOCK IMPLEMENTATIONS ---
	requests.AskMusicGenre = func(task string) string {
		return "lofi"
	}

	requests.RetrievePlaylist = func(genre string) requests.SpotifyPlaylistResponse {
		return requests.SpotifyPlaylistResponse{
			Playlists: requests.Playlists{
				Items: []requests.PlaylistItem{
					{
						Name:        "Lo-fi Beats",
						Description: "Chill and focus",
						ExternalURLs: requests.ExternalURLs{
							Spotify: "https://open.spotify.com/playlist/123",
						},
						Images: []requests.Image{
							{Url: "https://image.url/lofi.jpg"},
						},
					},
				},
			},
		}
	}

	storage.CreateTodo = func(todo *models.Todo) error {
		// Pretend it was saved and assign an ID
		todo.ID = 1
		return nil
	}

	// --- TEST REQUEST BODY ---
	inputTodo := models.Todo{Title: "Do the dishes"}
	body, _ := json.Marshal(inputTodo)

	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()

	// --- CALL HANDLER ---
	handlers.CreateTodo(w, req)

	// --- ASSERTIONS ---
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	var got models.Todo
	err := json.NewDecoder(resp.Body).Decode(&got)
	assert.NoError(t, err)

	assert.Equal(t, 1, got.ID)
	assert.Equal(t, "Do the dishes", got.Title)
	assert.Equal(t, "Lo-fi Beats", got.Playlist.Name)
	assert.Equal(t, "Chill and focus", got.Playlist.Description)
	assert.Equal(t, "https://open.spotify.com/playlist/123", got.Playlist.Link)
	assert.Equal(t, "https://image.url/lofi.jpg", got.Playlist.Image)
}

func TestCreateTodo_InvalidInput(t *testing.T) {
	req := httptest.NewRequest("POST", "/todos", bytes.NewReader([]byte("{invalid json")))
	w := httptest.NewRecorder()

	handlers.CreateTodo(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
