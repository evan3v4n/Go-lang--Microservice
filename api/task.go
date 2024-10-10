// Path: project_root/api/models/task.go

package models

import (
	"time"
)

type Task struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" binding:"required,min=3,max=100"`
	Description string    `json:"description" binding:"max=500"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
