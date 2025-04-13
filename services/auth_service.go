package services

import (
	"Task_manager_apis/models"
	"Task_manager_apis/utils"
	"errors"
)

// AuthService struct to group authentication function
type AuthService struct {
	UserRepo *models.UserRepository
}

// NewAuthService creatres a new instance of AuthService
func NewAuthService(userRepo *models.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

// for CreateUser
func (s *AuthService) CreateUser(user *models.User) error {
	// validate user input
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	// Check if the user already exists
	exists, err := s.UserRepo.UserExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user with this email already exists")
	}

	// Hash the password before storing
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save user to the databse using the model function
	return s.UserRepo.CreateUser(user)
}

// Login handles user login and returns JWT if successful
func (s *AuthService) LoginUser(email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CompareHashPassword(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
