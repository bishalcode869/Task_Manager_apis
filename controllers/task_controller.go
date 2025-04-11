package controllers

import (
	"Task_manager_apis/models"
	"Task_manager_apis/services"
	"Task_manager_apis/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Home returns a welcome message
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Task_Manager_Application"})
}

// AddTask adds a new task to the database
func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Get user ID from context
	userID := c.GetUint("userID")
	newTask.UserID = userID

	service := services.NewTaskService()
	if err := service.CreateTask(&newTask); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to add task")
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// ListTasks retrieves a list of tasks with pagination
func ListTasks(c *gin.Context) {
	userID := c.GetUint("userID")

	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid page number")
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid limit number")
		return
	}

	service := services.NewTaskService()
	tasks, err := service.GetAllTasks(userID, page, limit)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetByID retrieves a task by its ID
func GetByID(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	service := services.NewTaskService()
	task, err := service.GetTaskByID(id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.HandleError(c, http.StatusNotFound, "Task not found")
		} else {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to retrieve task")
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

// MarkTaskDone marks a task as done
func MarkTaskDone(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	service := services.NewTaskService()
	task, err := service.MarkTaskDone(id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.HandleError(c, http.StatusNotFound, "Task not found")
		} else {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to update task status")
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by its ID
func DeleteTask(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	service := services.NewTaskService()
	if err := service.DeleteTask(id, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.HandleError(c, http.StatusNotFound, "Task not found")
		} else {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to delete task")
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task successfully deleted"})
}
