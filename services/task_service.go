package services

import (
	"Task_manager_apis/models"
	"errors"
)

type TaskService struct {
	TaskRepo *models.TaskRepository
}

// create new instance
func NewTaskService(taskRepo *models.TaskRepository) *TaskService {
	return &TaskService{TaskRepo: taskRepo}
}

// CreateTask adds a new task
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.TaskRepo.CreateTask(task)
}

// GetAllTasks fetches all tasks with pagination for a user
func (s *TaskService) GetAllTasks(userID uint, pageNum, limitNum int) ([]models.Task, error) {
	return s.TaskRepo.GetAllTasks(int(userID), pageNum, limitNum)
}

// GetTaskByID retrieves a task by ID and user
func (s *TaskService) GetTaskByID(id int, userID uint) (*models.Task, error) {
	if id <= 0 {
		return nil, errors.New("invalid task ID")
	}
	return s.TaskRepo.GetTaskByID(id, userID)
}

// MarkTaskDone marks a task as done for a user
func (s *TaskService) MarkTaskDone(id int, userID uint) (*models.Task, error) {
	return s.TaskRepo.MarkTaskDone(id, userID)
}

// DeleteTask deletes a task for a user
func (s *TaskService) DeleteTask(id int, userID uint) error {
	return s.TaskRepo.DeleteTask(id, userID)
}
