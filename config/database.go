package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	LoadEnv()

	// Required environment variables for the database connection.
	requiredEnvVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	for _, envVar := range requiredEnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatalf("Error: Missing required environment variable %s", envVar)
		}
	}

	// Prepare the Data soure Name (DNS) for connection tothe PostgreSQL database.
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Try to open a connection to the database
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		// Log the error with details and stop execution if the connection fails.
		log.Fatal("Falied to connect to database: ", err)
	}

	// Set the global DB variable
	DB = db
	log.Println("successfully connected to the database")
}
