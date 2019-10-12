package model

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type (
	User struct {
		gorm.Model
		Name        string `gorm:"type:varchar(255);not_null"`
		Balance     float64
		Account     string `gorm:"type:varchar(36);not_null"`
		Skill       postgres.Jsonb
		Position    string
		Department  string `gorm:"type:varchar(32);not_null"`
		Achievement []Achievement
		Tasks       []Task
	}
	Task struct {
		gorm.Model
		Name           string `gorm:"type:varchar(255);not_null"`
		Score          int    `gorm:"not_null"`
		ExpirationDate time.Time
		ManagerID      uint
		UserID         uint
		TaskPriorityID uint
		TaskPriority   TaskPriority
		TaskStatusID   uint
		TaskStatus     TaskStatus
	}
	TaskPriority struct {
		gorm.Model
		Priority string `gorm:"type:varchar(16);not_null"`
	}
	TaskStatus struct {
		gorm.Model
		Status string `gorm:"type:varchar(16);not_null"`
	}
	Course struct {
	}
	Achievement struct {
		gorm.Model
		Name           string `gorm:"type:varchar(255);not_null" json:"name"`
		ExpirationDate time.Time
		UserID         uint
	}
)
