package transaction

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/transaction/dto"
	"ulinan/entities"
)

type TransactionRepositoryInterface interface {
	FindTransactions(page, perPage int) ([]*entities.OrderEntity, error)
	CountTotalTransactions() (int64, error)
	FindByName(page, perPage int, name string) ([]*entities.OrderEntity, error)
	CountTransactionByName(name string) (int64, error)
	FindTransactionById(id string) (*entities.OrderEntity, error)
	UpdatePaymentStatus(orderID string, paymentStatus string) error
	DeleteTransaction(orderID string) error
}

type TransactionServiceInterface interface {
	CalculatePaginationValues(page int, totalItems int, perPage int) (int, int)
	GetNextPage(currentPage, totalPages int) int
	GetPrevPage(currentPage int) int
	GetAllTransactions(page, perPage int) ([]*entities.OrderEntity, int64, error)
	GetTransactionById(id string) (*entities.OrderEntity, error)
	GetTransactionByUsername(page, perPage int, name string) ([]*entities.OrderEntity, int64, error)
	UpdatePaymentStatus(orderID string, request *dto.TUpdateTransactionRequest) error
	DeleteTransaction(orderID string) error
}

type TransactionHandlerInterface interface {
	GetAllTransactions(c *fiber.Ctx) error
	GetTransactionById(c *fiber.Ctx) error
	UpdatePaymentStatus(c *fiber.Ctx) error
	DeleteTransaction(c *fiber.Ctx) error
}
