package models

import (
	"Task_manager_apis/config"
	"errors"

	"gorm.io/gorm"
)

// Task represents the task model in the database
type Task struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Done   bool   `json:"done"`
	UserID uint   `json:"user_id"` // Foreign key to associate with a user
}

// CreateTask adds a new task to the database
func CreateTask(task *Task) error {
	return config.DB.Create(task).Error
}

// GetAllTasks retrieves all tasks from the database for a specific user with pagination
func GetAllTasks(userID, pageNum, limitNum int) ([]Task, error) {
	var tasks []Task
	err := config.DB.
		Where("user_id = ?", userID).
		Offset((pageNum - 1) * limitNum).
		Limit(limitNum).
		Find(&tasks).Error
	return tasks, err
}

// GetTaskByID retrieves a task by its ID and user
func GetTaskByID(id int, userID uint) (*Task, error) {
	var task Task
	err := config.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("task not found")
	}
	return &task, err
}

// MarkTaskDone updates the Done status of a task for a user
func MarkTaskDone(id int, userID uint) (*Task, error) {
	task, err := GetTaskByID(id, userID)
	if err != nil {
		return nil, err
	}

	task.Done = true
	if err := config.DB.Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask deletes a task by its ID and user
func DeleteTask(id int, userID uint) error {
	task, err := GetTaskByID(id, userID)
	if err != nil {
		return err
	}

	return config.DB.Delete(task).Error
}
