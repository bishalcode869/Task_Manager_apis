package services

import (
	"Task_manager_apis/models"
	"Task_manager_apis/utils"
	"errors"
	"fmt"
)

// AuthService struct to group authentication function
type AuthService struct{}

// NewAuthService creatres a new instance of AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// for CreateUser
func (s *AuthService) CreateUser(user *models.User) error {
	// validate user input
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	// Hash the password before storing
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save user to the databse using the model function
	return models.CreateUser(user)
}

// Login handles user login and returns JWT if successful
func (s *AuthService) LoginUser(email, password string) (string, error) {
	fmt.Println("Login attempt for:", email)

	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	fmt.Println("Found user:", user.Email)
	fmt.Println("Entered password:", password)
	fmt.Println("Stored hash:", user.Password)

	if !utils.CompareHashPassword(password, user.Password) {
		return "", errors.New("Invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}
	return token, nil

}
