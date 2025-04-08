package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Task struct defines a task model
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task   //Stores all tasks
var nextId int = 1 // Auto-increment ID

func main() {

	router := gin.Default() // Create a new Gin router

	// Home route - welcome message
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task manager web apis"})
	})

	// Create a new task
	router.POST("/tasks", func(ctx *gin.Context) {
		var newTask Task
		if err := ctx.BindJSON(&newTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid JSON"})
			return
		}
		newTask.Done = false
		newTask.ID = nextId
		nextId++
		tasks = append(tasks, newTask)
		ctx.JSON(http.StatusCreated, newTask)

	})

	// Get all tasks
	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, tasks)
	})

	// Get a task by ID
	router.GET("/tasks/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var id int
		if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}
		for _, task := range tasks {
			if task.ID == id {
				ctx.JSON(http.StatusOK, task)
				return
			}

		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	// Mark task as done
	router.PUT("/tasks/:id/done", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var id int
		if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}
		for i := range tasks {
			if tasks[i].ID == id {
				if tasks[i].Done {
					ctx.JSON(http.StatusOK, gin.H{"message": "Alread marked!"})
					return
				}
				tasks[i].Done = true
				ctx.JSON(http.StatusOK, tasks[i])
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})

	})

	// Delete a task by ID
	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var id int
		if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task Id"})
			return
		}
		index := -1
		for i, task := range tasks {
			if task.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		tasks = append(tasks[:index], tasks[index+1:]...)
		ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})

	router.Run(":8080") //Run the server on port 8080

}
