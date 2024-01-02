package user

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/user/dto"
	"ulinan/entities"
)

type UserRepositoryInterface interface {
	FindEmail(email string) (*entities.UserEntity, error)
	FindId(id int) (*entities.UserEntity, error)
	FindUsername(username string) (*entities.UserEntity, error)
	UpdateUserAvatar(userID int, avatarPath string) error
	UpdateUser(userUpdate *entities.UserEntity) error
	UpdateGender(user *entities.UserEntity) (*entities.UserEntity, error)
	UpdatePassword(userUpdate *entities.UserEntity) error
}

type UserServiceInterface interface {
	GetEmail(email string) (*entities.UserEntity, error)
	GetId(id int) (*entities.UserEntity, error)
	GetUsername(username string) (*entities.UserEntity, error)
	UpdateAvatar(userID int, request *dto.TUpdateAvatarRequest) (*entities.UserEntity, error)
	UpdateUser(userID int, request *dto.TUpdateUserRequest) (*entities.UserEntity, error)
	UpdatePassword(userID int, request *dto.TUpdatePasswordRequest) error
}

type UserHandlerInterface interface {
	GetCurrentUser(c *fiber.Ctx) error
	UpdateAvatar(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
}
