package controllers

import (
	"Task_manager_apis/config"
	"Task_manager_apis/models"
	"Task_manager_apis/services"
	"Task_manager_apis/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Create the UserRepository and AuthService
	userRepo := models.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)

	if err := authService.CreateUser(&newUser); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to register user: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":    newUser.ID,
			"email": newUser.Email,
		},
	})
}

func UserLogin(c *gin.Context) {
	var loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Create the UserRepository and AuthService
	userRepo := models.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)

	token, err := authService.LoginUser(loginCredentials.Email, loginCredentials.Password)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
