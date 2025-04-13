package controllers

import (
	"Task_manager_apis/config"
	"Task_manager_apis/models"
	"Task_manager_apis/services"
	"Task_manager_apis/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Task_Manager_Application"})
}

func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	userID := c.GetUint("userID")
	newTask.UserID = userID

	taskRepo := models.NewTaskRepository(config.DB)
	taskService := services.NewTaskService(taskRepo)

	if err := taskService.CreateTask(&newTask); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to add task")
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

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

	taskRepo := models.NewTaskRepository(config.DB)
	taskService := services.NewTaskService(taskRepo)

	tasks, err := taskService.GetAllTasks(userID, page, limit)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetByID(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	taskRepo := models.NewTaskRepository(config.DB)
	taskService := services.NewTaskService(taskRepo)

	task, err := taskService.GetTaskByID(id, userID)
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

func MarkTaskDone(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	taskRepo := models.NewTaskRepository(dbInstance.GetDB())
	taskService := services.NewTaskService(taskRepo)

	task, err := taskService.MarkTaskDone(id, userID)
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

func DeleteTask(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		utils.HandleError(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	taskRepo := models.NewTaskRepository(config.DB)
	taskService := services.NewTaskService(taskRepo)

	if err := taskService.DeleteTask(id, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.HandleError(c, http.StatusNotFound, "Task not found")
		} else {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to delete task")
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task successfully deleted"})
}
