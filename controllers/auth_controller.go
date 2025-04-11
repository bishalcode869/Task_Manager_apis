package controllers

import (
	"Task_manager_apis/models"
	"Task_manager_apis/services"
	"Task_manager_apis/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// function for /register
func UserRegister(c *gin.Context) {
	var newUser models.User

	// Bind JSON to the User model
	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Invalid input"+err.Error())
		return
	}

	authService := services.NewAuthService()

	// Attempt to create the user
	if err := authService.CreateUser(&newUser); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to register user: "+err.Error())
		return
	}

	// Respond with success
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registerd successfully",
		"user": gin.H{
			"id":    newUser.ID,
			"email": newUser.Email,
		},
	})

}

// function for /login
func UserLogin(c *gin.Context) {
	var loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the incoming JSON to the credentials structure
	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// User the AuthService to handle the login process
	authService := services.NewAuthService()

	// Call the Login functi nto authenticate the user and generate a jWT
	token, err := authService.LoginUser(loginCredentials.Email, loginCredentials.Password)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Respond with JWT token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
