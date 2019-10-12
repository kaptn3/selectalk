package model

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type (
	User struct {
		gorm.Model
		Balance     float64        `json:"balance"`
		Account     string         `gorm:"type:varchar(36);not_null" json:"account"`
		Skill       postgres.Jsonb `json:"skill"`
		Position    string         `json:"position"`
		Department  string         `gorm:"type:varchar(32);unique;not_null;unique_index" json:"department"`
		Achievement []Achievement
		Tasks       []Task
	}
	Task struct {
		gorm.Model
		Name           string `gorm:"type:varchar(255);not_null"`
		Score          int    `gorm:"not_null"`
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
		Name string `gorm:"type:varchar(255);not_null" json:"name"`
		UserID uint
	}
)
