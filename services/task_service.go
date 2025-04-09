package services

import (
	"Task_manager_apis/models"
	"errors"
)

// CreateTaskService adds a new task
func CreateTaskService(task *models.Task) error {
	return models.CreateTask(task)
}

// GetAllTasksService fetches all tasks with pagination
func GetAllTasksService(pageNum, limitNum int) ([]models.Task, error) {
	return models.GetAllTasks(pageNum, limitNum)
}

// GetTaskByIDService retrieves a task by ID
func GetTaskByIDService(id int) (*models.Task, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}
	return models.GetTaskByID(id)
}

// MarkTaskDoneService marks a task as done
func MarkTaskDoneService(id int) (*models.Task, error) {
	return models.MarkTaskDone(id)
}

// DeleteTaskService deletes a task
func DeleteTaskService(id int) error {
	return models.DeleteTask(id)
}
