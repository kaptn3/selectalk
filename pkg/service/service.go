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

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService() *TaskService {
	return &TaskService{
		repo:repository.NewTaskRepository(),
	}
}

func (s *TaskService) GetByUserID(id int) ([]model.Task, error) {
	tasks, err := s.repo.FindByUserID(id)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}
