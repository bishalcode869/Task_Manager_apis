package models

import (
	"Task_manager_apis/config"
	"fmt"

	"gorm.io/gorm"
)

// creating model
type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

// CreateUser
func CreateUser(user *User) error {
	fmt.Println("Saving user:", user.Email)
	return config.DB.Create(user).Error
}

// GetUserByEmail
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
