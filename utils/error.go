package utils

import "github.com/gin-gonic/gin"

// handleError simplifies error response handling
func HandleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}


