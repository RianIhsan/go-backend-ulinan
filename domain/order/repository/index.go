package repository

import (
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
	"ulinan/domain/order"
	"ulinan/domain/order/dto"
	"ulinan/entities"
	"ulinan/helper/payment"
)

type OrderRepository struct {
	db         *gorm.DB
	coreClient coreapi.Client
}

func NewOrderRepository(db *gorm.DB, coreClient coreapi.Client) order.OrderRepositoryInterface {
	return &OrderRepository{
		db:         db,
		coreClient: coreClient,
	}
}

func (r *OrderRepository) FindByName(page, perPage int, name string) ([]*entities.OrderEntity, error) {
	var orders []*entities.OrderEntity
	offset := (page - 1) * perPage
	query := r.db.Preload("User").Offset(offset).Limit(perPage).
		Joins("JOIN users ON users.id = orders.user_id").
		Where("orders.deleted_at IS NULL")
	if name != "" {
		query = query.Where("users.username LIKE ?", "%"+name+"%")
	}
	err := query.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) CreateOrder(newOrder *entities.OrderEntity) (*entities.OrderEntity, error) {
	newOrder.Id = uuid.New().String()
	err := r.db.Create(newOrder).Error
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (r *OrderRepository) ProcessGatewayPayment(totalAmountPaid int, orderID string, paymentMethod, name, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error) {
	var paymentType coreapi.CoreapiPaymentType

	switch paymentMethod {
	case "qris":
		paymentType = coreapi.PaymentTypeQris
	case "bri", "bni", "bca", "cimb", "permata":
		paymentType = coreapi.PaymentTypeBankTransfer
	case "gopay":
		paymentType = coreapi.PaymentTypeGopay
	}
	coreClient := r.coreClient
	resp, err := payment.CreateCoreAPIPaymentRequest(coreClient, orderID, totalAmountPaid, paymentType, name, email, phone, bank)
	if err != nil {
		log.Error(err)
		return nil, errors.New("failed to make a payment request")
	}
	return resp, nil
}

func (r *OrderRepository) GetOrderById(orderID string) (*entities.OrderEntity, error) {
	var orders entities.OrderEntity
	if err := r.db.
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Preload("OrderDetails.Product.ProductPhotos").
		Preload("User").
		Where("id = ? AND deleted_at IS NULL", orderID).
		First(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

func (r *OrderRepository) CheckTransaction(orderID string) (dto.Status, error) {
	var status dto.Status
	transactionStatusResp, err := r.coreClient.CheckTransaction(orderID)
	if err != nil {
		return dto.Status{}, err
	} else {
		if transactionStatusResp != nil {
			status = payment.TransactionStatus(transactionStatusResp)
			return status, nil
		}
	}
	return dto.Status{}, err
}

func (r *OrderRepository) ConfirmPayment(orderID, orderStatus, paymentStatus string) error {
	var orders entities.OrderEntity
	if err := r.db.Model(&orders).Where("id = ?", orderID).Updates(map[string]interface{}{
		"order_status":   orderStatus,
		"payment_status": paymentStatus,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetAllOrdersByUserID(userID int) ([]*entities.OrderEntity, error) {
	var orders []*entities.OrderEntity
	if err := r.db.
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Preload("OrderDetails.Product.ProductPhotos").
		Preload("User").
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
