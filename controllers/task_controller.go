package controllers

import (
	"Task_manager_apis/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// function for add task
func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newTask.ID = models.NextID
	models.NextID++
	models.Tasks = append(models.Tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

// function for get task list
func ListTasks(c *gin.Context) {
	c.JSON(http.StatusOK, models.Tasks)
}

// function for get task by id
func TaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, t := range models.Tasks {
		if t.ID == id {
			c.JSON(http.StatusOK, models.Tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// function for mark Task done
func MarkTaskDone(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, t := range models.Tasks {
		if t.ID == id {
			if t.Done {
				c.JSON(http.StatusOK, gin.H{"message": "Already marked"})
				return
			}
			models.Tasks[i].Done = true
			c.JSON(http.StatusOK, models.Tasks[i])
			return

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// function for delete Task
func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, t := range models.Tasks {
		if t.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
