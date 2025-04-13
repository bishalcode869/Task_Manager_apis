package models

import (
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

// TaskRepository handles database operations for tasks
type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

// CreateTask adds a new task to the database
func (r *TaskRepository) CreateTask(task *Task) error {
	return r.DB.Create(task).Error
}

// GetAllTasks retrieves all tasks from the database for a specific user with pagination
func (r *TaskRepository) GetAllTasks(userID, pageNum, limitNum int) ([]Task, error) {
	var tasks []Task
	err := r.DB.
		Where("user_id = ?", userID).
		Offset((pageNum - 1) * limitNum).
		Limit(limitNum).
		Find(&tasks).Error
	return tasks, err
}

// GetTaskByID retrieves a task by its ID and user
func (r *TaskRepository) GetTaskByID(id int, userID uint) (*Task, error) {
	var task Task
	err := r.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("task not found")
	}
	return &task, err
}

// MarkTaskDone updates the Done status of a task for a user
func (r *TaskRepository) MarkTaskDone(id int, userID uint) (*Task, error) {
	task, err := r.GetTaskByID(id, userID)
	if err != nil {
		return nil, err
	}

	task.Done = true
	if err := r.DB.Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask deletes a task by its ID and user
func (r *TaskRepository) DeleteTask(id int, userID uint) error {
	task, err := r.GetTaskByID(id, userID)
	if err != nil {
		return err
	}
	return r.DB.Delete(task).Error
}
