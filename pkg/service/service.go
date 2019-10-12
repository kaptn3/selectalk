package service

import (
	"github.com/rodsher/selectel/pkg/repository"
	"github.com/rodsher/selectel/pkg/model"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

func (s UserService) GetByID(id int) (model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
