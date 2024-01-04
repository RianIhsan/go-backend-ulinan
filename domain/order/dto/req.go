package dto

type TCreateOrderRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`

	PaymentMethod string `json:"payment_method" validate:"required"`
}

type Status struct {
	PaymentStatus string
	OrderStatus   string
}

type SendNotificationPaymentRequest struct {
	PaymentStatus string `json:"payment_status"`
	OrderID       string `json:"order_id"`
	UserID        int    `json:"user_id"`
	UserName      string `json:"user_name"`
	Title         string `json:"title"`
	Body          string `json:"body"`
}
