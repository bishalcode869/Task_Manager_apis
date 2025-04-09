package controllers

import (
	"Task_manager_apis/models"
	"Task_manager_apis/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// handleError simplifies error response handling
func handleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// Home returns a welcome message
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Task_Manager_Application"})
}

// AddTask adds a new task to the database
func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	if err := services.CreateTaskService(&newTask); err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to add task")
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// ListTasks retrieves a list of tasks with pagination
func ListTasks(c *gin.Context) {
	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		handleError(c, http.StatusBadRequest, "Invalid page number")
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		handleError(c, http.StatusBadRequest, "Invalid limit number")
		return
	}

	tasks, err := services.GetAllTasksService(page, limit)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetByID retrieves a task by its ID
func GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		handleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := services.GetTaskByIDService(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleError(c, http.StatusNotFound, "Task not found")
		} else {
			handleError(c, http.StatusInternalServerError, "Failed to retrieve task")
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

// MarkTaskDone marks a task as done
func MarkTaskDone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		handleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := services.MarkTaskDoneService(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleError(c, http.StatusNotFound, "Task not found")
		} else {
			handleError(c, http.StatusInternalServerError, "Failed to update task status")
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by its ID
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		handleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if err := services.DeleteTaskService(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleError(c, http.StatusNotFound, "Task not found")
		} else {
			handleError(c, http.StatusInternalServerError, "Failed to delete task")
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task successfully deleted"})
}
