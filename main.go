package main

import (
	"Task_manager_apis/config"
	"Task_manager_apis/models"
	"Task_manager_apis/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// main intializes the server an sets up routes
func main() {
	config.Connect()
	if err := config.DB.AutoMigrate(&models.Task{}); err != nil {
		fmt.Println("Migration failed", err)
		return
	}

	r := gin.Default()

	// Register routes
	routes.TaskRoutes(r)

	fmt.Printf("🚀 Server is running at http://localhost:8080")
	r.Run(":8080")
}
