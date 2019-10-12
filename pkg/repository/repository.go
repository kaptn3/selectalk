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

type AchievementRepository struct {
	db *gorm.DB
}

func NewAchievementRepository() *AchievementRepository {
	return &AchievementRepository{
		db: database.GetInstance(),
	}
}

func (r *AchievementRepository) Find() ([]model.Achievement, error) {
	var achievements []model.Achievement
	res := r.db.Find(&achievements)
	if res.Error != nil {
		return achievements, res.Error
	}

	return achievements, nil
}

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{
		db: database.GetInstance(),
	}
}

func (r *CourseRepository) Find() ([]model.Course, error) {
	var courses []model.Course
	res := r.db.Find(&courses)
	if res.Error != nil {
		return courses, res.Error
	}

	return courses, nil
}
