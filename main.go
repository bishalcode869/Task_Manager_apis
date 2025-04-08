package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create the type for Task
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// store and  id should be increment
var tasks []Task
var nextId int = 1

// main function
func main() {
	// create server with gin framework with gin route
	router := gin.Default()

	// now create apis for task manager first home
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task manager web apis"})
	})

	// post task, add the task
	router.POST("/tasks", func(ctx *gin.Context) {
		// create a new task
		var newTask Task

		// let's bind the post json to new tasks
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
	// get task to fetch all the tasks
	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, tasks)
	})

	// GET task by id
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
	// put task for marked the done
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
	// delete task for delete the task
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
	// listeinig server
	router.Run(":8080")

}
