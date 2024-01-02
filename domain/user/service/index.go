package service

import (
	"errors"
	"ulinan/domain/user"
	"ulinan/entities"
)

type UserService struct {
	repo user.UserRepositoryInterface
}

func NewUserService(repo user.UserRepositoryInterface) user.UserServiceInterface {
	return &UserService{repo: repo}
}

func (s *UserService) GetEmail(email string) (*entities.UserEntity, error) {
	result, err := s.repo.FindEmail(email)
	if err != nil {
		return nil, errors.New("your email has been already")
	}
	return result, nil
}

func (s *UserService) GetId(id int) (*entities.UserEntity, error) {
	result, err := s.repo.FindId(id)
	if err != nil {
		return nil, errors.New("id not found")
	}
	return result, nil
}

func (s *UserService) GetUsername(username string) (*entities.UserEntity, error) {
	result, err := s.repo.FindUsername(username)
	if err != nil {
		return nil, errors.New("username not found")
	}
	return result, nil
}
