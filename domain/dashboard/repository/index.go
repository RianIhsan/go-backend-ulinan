package repository

import (
	"gorm.io/gorm"
	"ulinan/domain/dashboard"
	"ulinan/entities"
)

type DashboardRepository struct {
	db *gorm.DB
}

func (r *DashboardRepository) CountAllCategories() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.CategoryEntity{}).Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *DashboardRepository) CountAllProducts() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.ProductEntity{}).Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *DashboardRepository) CountAllOrders() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.OrderEntity{}).Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *DashboardRepository) CountAllPaymentSuccess() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.OrderEntity{}).Where("payment_status = ? AND deleted_at IS NULL", "Success").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func NewDashboardRepository(db *gorm.DB) dashboard.DashboardRepositoryInterface {
	return &DashboardRepository{
		db: db,
	}
}
