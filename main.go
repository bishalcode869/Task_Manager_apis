package main

import (
	"Task_manager_apis/config"
	"Task_manager_apis/models"
	"Task_manager_apis/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// main intializes the server an sets up routes
func main() {
	// Connect to DB and get instance
	dbInstance, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	db := dbInstance.GetDB()

	// Run migrations (optional)
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal("Error running migrations: %v", err)
	}

	router := gin.Default()

	// Setup routes
	routes.UserRoutes(router)
	routes.TaskRoutes(router)

	fmt.Printf("🚀 Server is running at http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}
