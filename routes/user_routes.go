package routes

import (
	"Task_manager_apis/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.UserRegister)
	r.POST("/login", controllers.UserLogin)
}
