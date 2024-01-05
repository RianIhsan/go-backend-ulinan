package dashboard

import "github.com/gofiber/fiber/v2"

type DashboardRepositoryInterface interface {
	CountAllCategories() (int64, error)
	CountAllProducts() (int64, error)
	CountAllOrders() (int64, error)
	CountAllPaymentSuccess() (int64, error)
}

type DashboardServiceInterface interface {
	GetCardDashboard() (int64, int64, int64, int64, error)
}

type DashboardHandlerInterface interface {
	GetCardDashboard(c *fiber.Ctx) error
}
