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
	UpdateGender(user *entities.UserEntity) error
	UpdatePassword(userUpdate *entities.UserEntity) error
	FindGenderByID(genderID int) (*entities.GenderEntity, error)
	UpdateUserWithTransaction(user *entities.UserEntity) (*entities.UserEntity, error)
	UpdateUserGenderWithTransaction(user *entities.UserEntity, genderID int) error
}

type UserServiceInterface interface {
	GetEmail(email string) (*entities.UserEntity, error)
	GetId(id int) (*entities.UserEntity, error)
	GetUsername(username string) (*entities.UserEntity, error)
	UpdateAvatar(userID int, request *dto.TUpdateAvatarRequest) (*entities.UserEntity, error)
	UpdateUser(userID int, request *dto.TUpdateUserRequest) (*entities.UserEntity, error)
	UpdatePassword(userID int, request *dto.TUpdatePasswordRequest) error
	UpdateUserGender(userID int, genderID int) error
}

type UserHandlerInterface interface {
	GetCurrentUser(c *fiber.Ctx) error
	UpdateAvatar(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
}
