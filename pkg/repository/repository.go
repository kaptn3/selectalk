package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rodsher/selectel/pkg/model"
	"github.com/rodsher/selectel/pkg/database"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetInstance(),
	}
}

func (r *UserRepository) FindByID(id int) (model.User, error) {
	var user model.User
	res := r.db.Where("id = ?", id).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		db: database.GetInstance(),
	}
}

func (r TaskRepository) FindByUserID(userID int) ([]model.Task, error) {
	var tasks []model.Task
	res := r.db.Where("user_id = ?", userID).Find(&tasks)
	if res.Error != nil {
		return tasks, res.Error
	}

	return tasks, nil
}
