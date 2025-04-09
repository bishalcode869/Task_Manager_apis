package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model        // Includes fields lik ID, CreatedAt, UpdatedAt, and DeletedAt
	Title      string `json:"title"`
	Done       bool   `json:"done"`
}
