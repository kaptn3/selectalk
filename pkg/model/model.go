package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
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
		gorm.Model
		Name  string `gorm:"type:varchar(255);not_null"`
		Award float64
		Kind string `gorm:"type:varchar(255);not_null"`
	}
	Achievement struct {
		gorm.Model
		Name           string `gorm:"type:varchar(255);not_null"`
		ExpirationDate time.Time
		Cost           float64
	}
)
