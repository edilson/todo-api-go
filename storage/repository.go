package storage

import (
	"todo-api/config"
	"todo-api/models"
)

func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := config.DB.Find(&todos)
	return todos, result.Error
}

func GetTodoByID(id uint) (models.Todo, error) {
	var todo models.Todo
	result := config.DB.First(&todo, id)
	return todo, result.Error
}

func CreateTodo(todo *models.Todo) error {
	if !todo.Completed {
		todo.Completed = false
	}

	result := config.DB.Create(todo)
	return result.Error
}

func UpdateTodo(id uint, updatedData *models.Todo) (models.Todo, error) {
	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		return todo, err
	}

	todo.Title = updatedData.Title
	todo.Description = updatedData.Description
	todo.Completed = updatedData.Completed

	err := config.DB.Save(&todo).Error

	return todo, err
}

func DeleteTodo(id uint) error {
	result := config.DB.Delete(&models.Todo{}, id)
	return result.Error
}
