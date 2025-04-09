package controllers

import (
	"Task_manager_apis/config"
	"Task_manager_apis/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handleError simplifies error response handling
func handleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// Home function: returns a welcome message
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Welcome to Task_Manager_Application"})
}

// AddTask function: adds a new task to the database
func AddTask(c *gin.Context) {
	var newTask models.Task

	// Bind the input JSON to the new task struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create the new task in the database
	if err := config.DB.Create(&newTask).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to add task")
		return
	}

	// Return the created task
	c.JSON(http.StatusCreated, newTask)
}

// ListTasks function: retrieves a list of tasks with pagination
func ListTasks(c *gin.Context) {
	page := c.DefaultQuery("page", "1")    // Default page is 1
	limit := c.DefaultQuery("limit", "10") // Default limit is 10
	pageNum, _ := strconv.Atoi(page)
	limitNum, _ := strconv.Atoi(limit)

	// Declare a slice to hold the tasks
	var tasks []models.Task

	// Fetch the tasks with pagination
	if err := config.DB.Offset((pageNum - 1) * limitNum).Limit(limitNum).Find(&tasks).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	// Return the list of tasks
	c.JSON(http.StatusOK, tasks)
}

// GetByID function: retrieves a task by its ID
func GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input ID")
		return
	}

	var task models.Task
	// Fetch the task by its ID
	if err := config.DB.First(&task, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "Task not found")
		return
	}

	// Return the task
	c.JSON(http.StatusOK, task)
}

// MarkTaskDone function: marks a task as done
func MarkTaskDone(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input ID")
		return
	}

	var task models.Task
	// Fetch the task by its ID
	if err := config.DB.First(&task, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "Task not found")
		return
	}

	// Check if the task is already marked as done
	if task.Done {
		c.JSON(http.StatusOK, gin.H{"message": "Task is already marked as done"})
		return
	}

	// Mark the task as done
	task.Done = true

	// Update the task in the database
	if err := config.DB.Save(&task).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to update task")
		return
	}

	// Return the updated task
	c.JSON(http.StatusOK, task)
}

// DeleteTask function: deletes a task by its ID
func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input ID")
		return
	}

	var task models.Task
	// Fetch the task by its ID
	if err := config.DB.First(&task, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "Task not found")
		return
	}

	// Delete the task from the database
	if err := config.DB.Delete(&task).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Task successfully deleted"})
}
