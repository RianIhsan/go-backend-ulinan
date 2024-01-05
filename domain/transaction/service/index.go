package service

import (
	"errors"
	"math"
	"ulinan/domain/transaction"
	"ulinan/domain/transaction/dto"
	"ulinan/entities"
)

type TransactionService struct {
	repo transaction.TransactionRepositoryInterface
}

func NewTransactionService(repo transaction.TransactionRepositoryInterface) transaction.TransactionServiceInterface {
	return &TransactionService{repo}
}

func (s *TransactionService) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
	pageInt := page
	if pageInt <= 0 {
		pageInt = 1
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if pageInt > totalPages {
		pageInt = totalPages
	}

	return pageInt, totalPages
}

func (s *TransactionService) GetNextPage(currentPage, totalPages int) int {
	if currentPage < totalPages {
		return currentPage + 1
	}
	return totalPages
}

func (s *TransactionService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}

func (s *TransactionService) GetAllTransactions(page, perPage int) ([]*entities.OrderEntity, int64, error) {
	transactions, err := s.repo.FindTransactions(page, perPage)
	if err != nil {
		return nil, 0, err
	}
	totalTransactions, err := s.repo.CountTotalTransactions()
	if err != nil {
		return nil, 0, err
	}
	return transactions, totalTransactions, nil
}

func (s *TransactionService) GetTransactionByUsername(page, perPage int, name string) ([]*entities.OrderEntity, int64, error) {
	transactionList, err := s.repo.FindByName(page, perPage, name)
	if err != nil {
		return nil, 0, err
	}
	totalTransactions, err := s.repo.CountTransactionByName(name)
	if err != nil {
		return nil, 0, err
	}
	return transactionList, totalTransactions, nil
}

func (s *TransactionService) GetTransactionById(id string) (*entities.OrderEntity, error) {
	transaction, err := s.repo.FindTransactionById(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) UpdatePaymentStatus(orderID string, request *dto.TUpdateTransactionRequest) error {
	transaction, err := s.repo.FindTransactionById(orderID)
	if err != nil {
		return errors.New("transaction data not found")
	}

	err = s.repo.UpdatePaymentStatus(transaction.Id, request.PaymentStatus)
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) DeleteTransaction(orderID string) error {
	transaction, err := s.repo.FindTransactionById(orderID)
	if err != nil {
		return errors.New("transaction data not found")
	}

	err = s.repo.DeleteTransaction(transaction.Id)
	if err != nil {
		return err
	}
	return nil
}
