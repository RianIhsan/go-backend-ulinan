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

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	TotalItems  int `json:"total_items"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
}

type PaginationRes struct {
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

func GetCurrentUser(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func SendStatusOkResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusCreatedResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusCreated).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusCreatedWithDataResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(GeneralMessageWithData{
		Message: message,
		Data:    data,
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

func SendPaginationResponse(c *fiber.Ctx, data interface{}, currentPage, totalPages, totalItems, nextPage, prevPage int, message string) error {
	pagination := PaginationMeta{
		CurrentPage: currentPage,
		TotalPage:   totalPages,
		TotalItems:  totalItems,
		NextPage:    nextPage,
		PrevPage:    prevPage,
	}
	return c.Status(fiber.StatusOK).JSON(PaginationRes{
		Message: message,
		Data:    data,
		Meta:    pagination,
	})
}
