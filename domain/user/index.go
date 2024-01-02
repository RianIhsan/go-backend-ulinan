package user

import "ulinan/entities"

type UserRepositoryInterface interface {
	FindEmail(email string) (*entities.UserEntity, error)
	FindId(id int) (*entities.UserEntity, error)
	FindUsername(username string) (*entities.UserEntity, error)
}

type UserServiceInterface interface {
	GetEmail(email string) (*entities.UserEntity, error)
	GetId(id int) (*entities.UserEntity, error)
	GetUsername(username string) (*entities.UserEntity, error)
}
