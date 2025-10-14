package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/config"
	"todo-api/handlers"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	config.ConnectDB()
	handler := handlers.RegisterHandler()

	body := []byte(`{"username":"john","password":"secret"}`)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "User registered successfully")
}

func TestLoginHandler(t *testing.T) {
	config.ConnectDB()
	jwtKey := []byte("test_secret")

	reqReg := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(`{"username":"john","password":"secret"}`))
	rrReg := httptest.NewRecorder()
	handlers.RegisterHandler().ServeHTTP(rrReg, reqReg)

	loginBody := []byte(`{"username":"john","password":"secret"}`)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handlers.LoginHandler(jwtKey).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "token")
}
