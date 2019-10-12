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
	result := r.db.Where("id = ?", id).Take(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
