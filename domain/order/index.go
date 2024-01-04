package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"ulinan/domain/order/dto"
	"ulinan/entities"
)

type OrderRepositoryInterface interface {
	FindByName(page, perPage int, name string) ([]*entities.OrderEntity, error)
	CreateOrder(newOrder *entities.OrderEntity) (*entities.OrderEntity, error)
	ProcessGatewayPayment(totalAmountPaid int, orderID string, paymentMethod, name, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error)
	GetOrderById(orderID string) (*entities.OrderEntity, error)
	CheckTransaction(orderID string) (dto.Status, error)
	ConfirmPayment(orderID, orderStatus, paymentStatus string) error
}

type OrderServiceInterface interface {
	CreateOrder(userID int, request *dto.TCreateOrderRequest, bank midtrans.Bank) (interface{}, error)
	ProcessGatewayPayment(totalAmountPaid int, orderID string, paymentMethod, name, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error)
	ConfirmPayment(orderID string) error
	SendNotificationPayment(request dto.SendNotificationPaymentRequest) (string, error)
	CancelPayment(orderID string) error
	CallBack(notifPayload map[string]interface{}) error
}

type OrderHandlerInterface interface {
	CreateOrder(c *fiber.Ctx) error
	Callback(c *fiber.Ctx) error
}
