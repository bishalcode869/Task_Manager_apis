package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// User model represents users in the system
type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

// UserRepository provides methods for interacting with users
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser stores a new user
func (r *UserRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

// GetUserByEmail retrieves a user by their email
func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user with email %s not found", email)
	}
	return &user, err
}

// Optional: Check if user exists
func (r *UserRepository) UserExists(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
