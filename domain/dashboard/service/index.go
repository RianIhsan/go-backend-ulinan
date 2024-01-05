package service

import (
	"ulinan/domain/dashboard"
)

type DashboardService struct {
	repo dashboard.DashboardRepositoryInterface
}

func NewDashboardService(repo dashboard.DashboardRepositoryInterface) dashboard.DashboardServiceInterface {
	return &DashboardService{
		repo: repo,
	}
}

func (s *DashboardService) GetCardDashboard() (int64, int64, int64, int64, error) {
	productCount, err := s.repo.CountAllProducts()
	if err != nil {
		return 0, 0, 0, 0, err
	}
	categoryCount, err := s.repo.CountAllCategories()
	if err != nil {
		return 0, 0, 0, 0, err
	}
	orderCount, err := s.repo.CountAllOrders()
	if err != nil {
		return 0, 0, 0, 0, err
	}
	paymentSuccessCount, err := s.repo.CountAllPaymentSuccess()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return productCount, categoryCount, orderCount, paymentSuccessCount, nil
}
