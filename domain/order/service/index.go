package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"time"
	"ulinan/domain/cart"
	"ulinan/domain/order"
	"ulinan/domain/order/dto"
	"ulinan/domain/product"
	"ulinan/domain/user"
	"ulinan/entities"
	"ulinan/helper/generator"
)

type OrderService struct {
	repo           order.OrderRepositoryInterface
	generatorID    generator.GeneratorInterface
	productService product.ProductServiceInterface
	userService    user.UserServiceInterface
	cartService    cart.CartServiceInterface
}

func NewOrderService(
	repo order.OrderRepositoryInterface,
	generatorID generator.GeneratorInterface,
	productService product.ProductServiceInterface,
	userService user.UserServiceInterface,
	cartService cart.CartServiceInterface,
) order.OrderServiceInterface {
	return &OrderService{
		repo:           repo,
		generatorID:    generatorID,
		productService: productService,
		userService:    userService,
		cartService:    cartService,
	}
}

func (s *OrderService) CreateOrder(userID int, request *dto.TCreateOrderRequest, bank midtrans.Bank) (interface{}, error) {
	orderID, err := s.generatorID.GenerateOrderID()
	if err != nil {
		return nil, errors.New("failed create uuid order")
	}
	idOrder, err := s.generatorID.GenerateOrderID()
	if err != nil {
		return nil, errors.New("failed create id_order")
	}
	var validPaymentMethods = map[string]bool{
		"qris":          true,
		"bank_transfer": true,
		"gopay":         true,
		"bca":           true,
		"bri":           true,
		"bni":           true,
		"cimb":          true,
	}

	if !validPaymentMethods[request.PaymentMethod] {
		return nil, errors.New("invalid payment type")
	}

	var orderDetails []entities.OrderDetailsEntity
	var totalQuantity, totalPrice int

	products, err := s.productService.GetProductById(request.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	orderDetail := entities.OrderDetailsEntity{
		OrderId:    orderID,
		ProductId:  request.ProductID,
		Quantity:   request.Quantity,
		TotalPrice: request.Quantity * products.Price,
	}
	totalQuantity += request.Quantity
	totalPrice += orderDetail.TotalPrice

	orderDetails = append(orderDetails, orderDetail)

	if isInCart := s.cartService.IsProductInCart(userID, products.ID); isInCart {
		if err := s.cartService.RemoveProductFromCart(userID, products.ID); err != nil {
			return nil, errors.New("failed to delete cart")
		}
	}
	newData := &entities.OrderEntity{
		Id:                 orderID,
		IdOrder:            idOrder,
		UserId:             userID,
		GrandTotalQuantity: totalQuantity,
		GrandTotalPrice:    totalPrice,
		TotalAmountPaid:    totalPrice,
		OrderStatus:        "Menunggu Konfirmasi",
		PaymentStatus:      "Menunggu Konfirmasi",
		PaymentMethod:      request.PaymentMethod,
		CreatedAt:          time.Now(),
		OrderDetails:       orderDetails,
	}

	createdOrder, err := s.repo.CreateOrder(newData)
	if err != nil {
		return nil, errors.New("failed create order")
	}

	user, err := s.userService.GetId(createdOrder.UserId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	var phone string
	if user != nil && user.Phone != nil {
		phone = *user.Phone
	}
	switch request.PaymentMethod {
	case "qris", "bank_transfer", "bca", "bri", "bni", "cimb", "gopay":
		return s.ProcessGatewayPayment(totalPrice, createdOrder.Id, request.PaymentMethod, user.Username, user.Email, phone, bank)
	default:
		return nil, errors.New("invalid payment type")
	}
}

func (s *OrderService) ProcessGatewayPayment(totalAmountPaid int, orderID string, paymentMethod, name, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error) {
	result, err := s.repo.ProcessGatewayPayment(totalAmountPaid, orderID, paymentMethod, name, email, phone, bank)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *OrderService) ConfirmPayment(orderID string) error {
	orders, err := s.repo.GetOrderById(orderID)
	if err != nil {
		return errors.New("order not found")
	}
	orders.OrderStatus = "Proses"
	orders.PaymentStatus = "Konfirmasi"

	if err := s.repo.ConfirmPayment(orders.Id, orders.OrderStatus, orders.PaymentStatus); err != nil {
		return err
	}

	user, err := s.userService.GetId(orders.UserId)
	if err != nil {
		return errors.New("user not found")
	}

	notificationRequest := dto.SendNotificationPaymentRequest{
		OrderID:       orders.Id,
		UserID:        user.ID,
		PaymentStatus: "Konfirmasi",
	}
	_, err = s.SendNotificationPayment(notificationRequest)
	if err != nil {
		log.Error("failed to send notification: ", err)
		return err
	}
	return nil
}

func (s *OrderService) SendNotificationPayment(request dto.SendNotificationPaymentRequest) (string, error) {
	var notificationMsg string
	var err error

	user, err := s.userService.GetId(request.UserID)
	if err != nil {
		return "", err
	}
	orders, err := s.repo.GetOrderById(request.OrderID)
	if err != nil {
		return "", err
	}
	switch request.PaymentStatus {
	case "Menunggu Konfirmasi":
		notificationMsg = fmt.Sprintf("Alloo, %s! Pesananmu dengan ID %s udah berhasil dibuat, nih. Ditunggu yupp!!", user.Username, orders.IdOrder)
	case "Konfirmasi":
		notificationMsg = fmt.Sprintf("Thengkyuu, %s! Pembayaran untuk pesananmu dengan ID %s udah kami terima, nih. Semoga harimu menyenangkan!", user.Username, orders.IdOrder)
	case "Gagal":
		notificationMsg = fmt.Sprintf("Maaf, %s. Pembayaran untuk pesanan dengan ID %s gagal, nih. Beritahu kami apabila kamu butuh bantuan yaa!!", user.Username, orders.IdOrder)
	default:
		return "", errors.New("status pesanan tidak valid")
	}

	return notificationMsg, nil
}

func (s *OrderService) CancelPayment(orderID string) error {
	orders, err := s.repo.GetOrderById(orderID)
	if err != nil {
		return errors.New("order not found")
	}
	orders.OrderStatus = "Gagal"
	orders.PaymentStatus = "Gagal"

	if err := s.repo.ConfirmPayment(orders.Id, orders.OrderStatus, orders.PaymentStatus); err != nil {
		return errors.New("failed to cancel the order")
	}

	user, err := s.userService.GetId(orders.UserId)
	if err != nil {
		return errors.New("user not found")
	}

	notificationRequest := dto.SendNotificationPaymentRequest{
		OrderID:       orders.Id,
		UserID:        user.ID,
		PaymentStatus: "Gagal",
	}
	_, err = s.SendNotificationPayment(notificationRequest)
	if err != nil {
		log.Error("failed to send notification: ", err)
		return err
	}
	return nil
}

func (s *OrderService) CallBack(notifPayload map[string]interface{}) error {
	orderID, exist := notifPayload["order_id"].(string)
	if !exist {
		return errors.New("invalid notification payload")
	}

	status, err := s.repo.CheckTransaction(orderID)
	if err != nil {
		return err
	}

	transaction, err := s.repo.GetOrderById(orderID)
	if err != nil {
		return errors.New("transaction data not found")
	}

	if status.PaymentStatus == "Konfirmasi" {
		if err := s.ConfirmPayment(transaction.Id); err != nil {
			return err
		}
	} else if status.PaymentStatus == "Gagal" {
		if err := s.CancelPayment(transaction.Id); err != nil {
			return err
		}
	}

	return nil
}
