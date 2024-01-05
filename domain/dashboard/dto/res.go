package dto

type GetCardDashboardResponse struct {
	ProductCount        int64 `json:"product_count"`
	CategoryCount       int64 `json:"category_count"`
	OrderCount          int64 `json:"order_count"`
	PaymentSuccessCount int64 `json:"payment_success_count"`
}

func FormatGetCardDashboard(productCount, categoryCount, orderCount, paymentSuccessCount int64) *GetCardDashboardResponse {
	cardDashboardResponse := &GetCardDashboardResponse{
		ProductCount:        productCount,
		CategoryCount:       categoryCount,
		OrderCount:          orderCount,
		PaymentSuccessCount: paymentSuccessCount,
	}
	return cardDashboardResponse
}
