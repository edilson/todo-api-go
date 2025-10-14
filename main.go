package main

import (
	"log"
	"net/http"
	"todo-api/config"
	"todo-api/models"
	"todo-api/routes"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Todo{}, &models.User{})

	router := routes.SetupRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
