package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"mime/multipart"
	"ulinan/domain/user"
	"ulinan/domain/user/dto"
	"ulinan/entities"
	"ulinan/helper/cloudinary"
	"ulinan/helper/response"
)

type userHandler struct {
	userService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) user.UserHandlerInterface {
	return &userHandler{userService: userService}
}

func (u userHandler) GetCurrentUser(c *fiber.Ctx) error {
	user, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || user == nil {
		return response.SendStatusUnauthorized(c, "user not found")
	}

	return response.GetCurrentUser(c, dto.GetUserResponse(user))
}

func (u userHandler) UpdateAvatar(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || currentUser == nil {
		return response.SendStatusUnauthorized(c, "User not found")
	}

	file, err := c.FormFile("avatar")
	var uploadedURL string
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed open to open file: "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)
		uploadedURL, err = cloudinary.Uploader(fileToUpload)
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed to upload image: "+err.Error())
		}
	}

	userUpdateAvatar := &dto.TUpdateAvatarRequest{
		Avatar: uploadedURL,
	}

	image, err := u.userService.UpdateAvatar(currentUser.ID, userUpdateAvatar)
	if err != nil {
		return response.SendStatusBadRequest(c, "Error upload image: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Success updating avatar", dto.UpdateAvatarResponse(image))
}

func (u userHandler) UpdateUser(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || currentUser == nil {
		return response.SendStatusUnauthorized(c, "User not found")
	}

	var request *dto.TUpdateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return response.SendStatusBadRequest(c, "Invalid request format")
	}

	_, err := u.userService.UpdateUser(currentUser.ID, request)
	if err != nil {
		log.Fatal(err.Error())
		return response.SendStatusInternalServerError(c, "Failed to update user information")
	}

	return response.SendStatusOkResponse(c, "User information updated successfully")
}

func (u userHandler) UpdatePassword(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || currentUser == nil {
		return response.SendStatusUnauthorized(c, "User not found")
	}

	var request dto.TUpdatePasswordRequest
	if err := c.BodyParser(&request); err != nil {
		return response.SendStatusBadRequest(c, "Invalid request payload")
	}

	err := u.userService.UpdatePassword(currentUser.ID, &request)
	if err != nil {
		return response.SendStatusBadRequest(c, "Error updating password: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "Password updated successfully")
}
