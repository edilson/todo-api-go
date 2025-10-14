package routes

import (
	"net/http"
	"os"
	"todo-api/handlers"
	"todo-api/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	jwtKet := []byte(os.Getenv("AUTH_SECRET"))

	router.Handle("/todos", middlewares.JWTMiddleware(jwtKet)(http.HandlerFunc(handlers.GetTodos))).Methods("GET")
	router.Handle("/todos/{id}", middlewares.JWTMiddleware(jwtKet)(http.HandlerFunc(handlers.GetTodo))).Methods("GET")
	router.Handle("/todos", middlewares.JWTMiddleware(jwtKet)(http.HandlerFunc(handlers.CreateTodo))).Methods("POST")
	router.Handle("/todos/{id}", middlewares.JWTMiddleware(jwtKet)(http.HandlerFunc(handlers.UpdateTodo))).Methods("PUT")
	router.Handle("/todos/{id}", middlewares.JWTMiddleware(jwtKet)(http.HandlerFunc(handlers.DeleteTodo))).Methods("DELETE")

	router.HandleFunc("/login", handlers.LoginHandler(jwtKet)).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler()).Methods("POST")

	return router
}
