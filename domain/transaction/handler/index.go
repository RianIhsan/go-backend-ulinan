package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"ulinan/domain/transaction"
	"ulinan/domain/transaction/dto"
	"ulinan/entities"
	"ulinan/helper/response"
)

type TransactionHandler struct {
	service transaction.TransactionServiceInterface
}

func NewTransactionHandler(service transaction.TransactionServiceInterface) transaction.TransactionHandlerInterface {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageConv, _ := strconv.Atoi(strconv.Itoa(page))
	limit := c.Query("limit", "10")
	limitInt, _ := strconv.Atoi(limit)

	var transactions []*entities.OrderEntity
	var totalItems int64
	var err error
	search := c.Query("search")

	if search != "" {
		transactions, totalItems, err = h.service.GetTransactionByUsername(pageConv, limitInt, search)
	} else if limit != "" {
		transactions, totalItems, err = h.service.GetAllTransactions(pageConv, limitInt)
	}

	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to get transactions: "+err.Error())
	}

	currentPage, totalPages := h.service.CalculatePaginationValues(pageConv, int(totalItems), limitInt)
	nextPage := h.service.GetNextPage(currentPage, totalPages)
	prevPage := h.service.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.FormatterGetAllTransaction(transactions), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success get categories")
}

func (h *TransactionHandler) GetTransactionById(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	id := c.Params("id")

	transaction, err := h.service.GetTransactionById(id)
	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to get transaction: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success get transaction", dto.FormatGetTransactionById(transaction))
}

func (h *TransactionHandler) UpdatePaymentStatus(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	id := c.Params("id")

	var request dto.TUpdateTransactionRequest
	if err := c.BodyParser(&request); err != nil {
		return response.SendStatusBadRequest(c, "invalid request body")
	}

	if err := h.service.UpdatePaymentStatus(id, &request); err != nil {
		return response.SendStatusInternalServerError(c, "failed to update transaction: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "success update transaction")
}

func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	id := c.Params("id")

	if err := h.service.DeleteTransaction(id); err != nil {
		return response.SendStatusInternalServerError(c, "failed to delete transaction: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "success delete transaction")
}
