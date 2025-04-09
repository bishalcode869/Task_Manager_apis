package models

import (
	"Task_manager_apis/config" // Ensure correct import path for your config package

	"gorm.io/gorm"
)

// Task represents the task model in the database
type Task struct {
	gorm.Model        // Includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt
	Title      string `json:"title"`
	Done       bool   `json:"done"`
}

// CreateTask adds a new task to the database
func CreateTask(task *Task) error {
	return config.DB.Create(task).Error
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks(pageNum, limitNum int) ([]Task, error) {
	var tasks []Task
	err := config.DB.Offset((pageNum - 1) * limitNum).Limit(limitNum).Find(&tasks).Error
	return tasks, err
}

// GetTaskByID retrieves a task by its ID from the database
func GetTaskByID(id int) (*Task, error) {
	var task Task
	err := config.DB.First(&task, id).Error
	return &task, err
}

// MarkTaskDone updates the Done status of a task in the database
func MarkTaskDone(id int) (*Task, error) {
	var task Task
	err := config.DB.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	task.Done = true
	if err := config.DB.Save(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTask deletes a task by its ID from the database
func DeleteTask(id int) error {
	if err := config.DB.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
