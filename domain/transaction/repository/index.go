package repository

import (
	"gorm.io/gorm"
	"strings"
	"time"
	"ulinan/domain/transaction"
	"ulinan/entities"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepositoryInterface {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) FindTransactions(page, perPage int) ([]*entities.OrderEntity, error) {
	var transactions []*entities.OrderEntity
	offset := (page - 1) * perPage
	err := r.db.
		Offset(offset).
		Limit(perPage).
		Preload("User").
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Where("deleted_at IS NULL").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) CountTotalTransactions() (int64, error) {
	var count int64
	err := r.db.Model(&entities.OrderEntity{}).
		Preload("User").
		Where("deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TransactionRepository) FindByName(page, perPage int, name string) ([]*entities.OrderEntity, error) {
	var transactions []*entities.OrderEntity
	offset := (page - 1) * perPage
	query := r.db.Offset(offset).Limit(perPage).Where("orders.deleted_at IS NULL") // Menggunakan 'orders.deleted_at' untuk menghindari ambiguitas

	if name != "" {
		query = query.Preload("User").Joins("JOIN users ON users.id = orders.user_id").Where("LOWER(users.username) LIKE ?", "%"+strings.ToLower(name)+"%")
	}

	err := query.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) CountTransactionByName(name string) (int64, error) {
	var count int64
	query := r.db.Model(&entities.OrderEntity{}).Joins("JOIN users ON users.id = orders.user_id").Where("orders.deleted_at IS NULL") // Menggunakan 'orders.deleted_at' untuk menghindari ambiguitas

	if name != "" {
		query = query.Where("LOWER(users.username) LIKE ?", "%"+strings.ToLower(name)+"%")
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TransactionRepository) FindTransactionById(id string) (*entities.OrderEntity, error) {
	var transaction *entities.OrderEntity
	err := r.db.Preload("User").
		Preload("OrderDetails.Product").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) UpdatePaymentStatus(orderID string, paymentStatus string) error {
	updateFields := map[string]interface{}{
		"payment_status": paymentStatus,
		"updated_at":     time.Now(),
	}

	if err := r.db.Model(&entities.OrderEntity{}).Where("id = ? AND deleted_at IS NULL", orderID).Updates(updateFields).Error; err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) DeleteTransaction(orderID string) error {
	category := &entities.OrderEntity{}
	if err := r.db.Model(category).Where("id = ?", orderID).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
