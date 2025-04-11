package routes

import (
	"Task_manager_apis/controllers"
	"Task_manager_apis/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	r.GET("/", controllers.Home)
	tasks := r.Group("/tasks")
	{
		tasks.POST("/", middleware.JWTAuthMiddleware(), controllers.AddTask)
		tasks.GET("/", middleware.JWTAuthMiddleware(), controllers.ListTasks)
		tasks.GET("/:id", middleware.JWTAuthMiddleware(), controllers.GetByID)
		tasks.PUT("/:id/done", middleware.JWTAuthMiddleware(), controllers.MarkTaskDone)
		tasks.DELETE("/:id", middleware.JWTAuthMiddleware(), controllers.DeleteTask)
	}
}
