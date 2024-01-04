package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"ulinan/domain/order"
	"ulinan/domain/order/dto"
	"ulinan/entities"
	"ulinan/helper/response"
	"ulinan/helper/validator"
)

type OrderHandler struct {
	service order.OrderServiceInterface
}

func NewOrderHandler(service order.OrderServiceInterface) order.OrderHandlerInterface {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "user" {
		return response.SendStatusUnauthorized(c, "Access denied: you are admin, not user")
	}

	var payload dto.TCreateOrderRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}

	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	var bank midtrans.Bank
	switch payload.PaymentMethod {
	case "bca":
		bank = midtrans.BankBca
	// Tambahkan bank lain sesuai kebutuhan
	case "bri":
		bank = midtrans.BankBri
	case "bni":
		bank = midtrans.BankBni
	case "cimb":
		bank = midtrans.BankCimb
	default:
		bank = midtrans.BankPermata
	}

	result, err := h.service.CreateOrder(currentUser.ID, &payload, bank)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create order: "+err.Error())
	}

	switch payload.PaymentMethod {
	case "qris", "bank_transfer", "bca", "bri", "bni", "cimb", "gopay":
		return response.SendStatusCreatedWithDataResponse(c, "Berhasil membuat pesanan dengan payment gateway", result)
	default:
		return response.SendStatusBadRequest(c, "Metode pembayaran tidak valid: "+payload.PaymentMethod)
	}
}

func (h *OrderHandler) Callback(c *fiber.Ctx) error {
	var notificationPayload map[string]any

	// Parse request body using Fiber's BodyParser
	if err := c.BodyParser(&notificationPayload); err != nil {
		return response.SendStatusBadRequest(c, "Format input yang Anda masukkan tidak sesuai")
	}

	err := h.service.CallBack(notificationPayload)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Gagal callback pesanan: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "Berhasil callback")
}

func (h *OrderHandler) GetAllOrdersByUserID(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "user" {
		return response.SendStatusUnauthorized(c, "Access denied: you are admin, not user")
	}

	orders, err := h.service.GetAllOrdersByUserID(currentUser.ID)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Gagal mendapatkan pesanan: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Berhasil mendapatkan pesanan", dto.FormatterGetAllOrderUser(orders))
}
