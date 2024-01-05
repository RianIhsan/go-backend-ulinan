package handler

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/dashboard"
	"ulinan/domain/dashboard/dto"
	"ulinan/entities"
	"ulinan/helper/response"
)

type DashboardHandler struct {
	service dashboard.DashboardServiceInterface
}

func NewDashboardHandler(service dashboard.DashboardServiceInterface) dashboard.DashboardHandlerInterface {
	return &DashboardHandler{
		service: service,
	}
}

func (h *DashboardHandler) GetCardDashboard(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	productCount, categoryCount, orderCount, paymentSuccessCount, err := h.service.GetCardDashboard()
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to get card dashboard: "+err.Error())
	}
	return response.SendStatusCreatedWithDataResponse(c, "success get card dashboard", dto.FormatGetCardDashboard(productCount, categoryCount, orderCount, paymentSuccessCount))
}
