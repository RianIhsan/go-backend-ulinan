package response

import (
	"github.com/gofiber/fiber/v2"
)

type GeneralMessage struct {
	Message string `json:"message"`
}

type GeneralMessageWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetCurrentUser(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func SendStatusOkResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusOkWithDataResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(GeneralMessageWithData{
		Message: message,
		Data:    data,
	})
}

func SendStatusBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusNotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusInternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusUnauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusForbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(GeneralMessage{
		Message: message,
	})
}
