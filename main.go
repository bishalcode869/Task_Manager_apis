package main

import (
	"Task_manager_apis/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// main intializes the server an sets up routes
func main() {
	r := gin.Default()

	// Register routes
	routes.TaskRoutes(r)

	fmt.Printf("🚀 Server is running at http://localhost:8080")
	r.Run(":8080")
}
