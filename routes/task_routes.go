package routes

import (
	"Task_manager_apis/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	tasks := r.Group("/tasks")
	{
		tasks.POST("/", controllers.AddTask)
		tasks.GET("/", controllers.ListTasks)
		tasks.GET("/:id", controllers.GetByID)
		tasks.PUT("/:id/done", controllers.MarkTaskDone)
		tasks.DELETE("/:id", controllers.DeleteTask)
	}
}
