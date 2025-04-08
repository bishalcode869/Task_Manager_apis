package main

import (
	"Task_manager_apis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// create routes
	r := gin.Default()

	// router
	routes.TaskRoutes(r)

	// running server at 8080
	r.Run(":8080")
}
