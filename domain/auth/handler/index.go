package handler

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/auth"
	"ulinan/domain/auth/dto"
	"ulinan/helper/response"
	"ulinan/helper/validator"
)

type AuthHandler struct {
	authService auth.AuthServiceInterface
}

func NewAuthHandler(authService auth.AuthServiceInterface) auth.AuthHandlerInterface {
	return &AuthHandler{authService: authService}
}

func (h AuthHandler) Register(c *fiber.Ctx) error {
	var payload dto.TRegisterRequest

	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}

	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	_, err := h.authService.Register(&payload)
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}
	return response.SendStatusOkResponse(c, "register is successfully")
}
func (h AuthHandler) Login(c *fiber.Ctx) error {
	var payload dto.TLoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}

	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	userLogin, accessToken, err := h.authService.Login(&payload)
	if err != nil {
		if err.Error() == "user not found" {
			return response.SendStatusNotFound(c, "user not found")
		}
		return response.SendStatusUnauthorized(c, "incorrect password")
	}

	return response.SendStatusOkWithDataResponse(c, "login successfully", dto.LoginResponse(userLogin, accessToken))
}
