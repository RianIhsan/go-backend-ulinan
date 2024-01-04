package dto

import "ulinan/entities"

type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type GetAllOrderUserResponse struct {
	ID            string `json:"id"`
	IdOrder       string `json:"id_order"`
	UserID        int    `json:"user_id"`
	ProductName   string `json:"product_name"`
	ArrivalDate   string `json:"arrival_date"`
	TotalPrice    int    `json:"total_price"`
	TotalQuantity int    `json:"total_quantity"`
	PaymentStatus string `json:"payment_status"`
}

func FormatGetAllOrderUser(order *entities.OrderEntity) *GetAllOrderUserResponse {
	orderResponse := &GetAllOrderUserResponse{}
	orderResponse.ID = order.Id
	orderResponse.IdOrder = order.IdOrder
	orderResponse.UserID = order.UserId
	orderResponse.ArrivalDate = order.ArrivalDate.Format("02 January 2006")
	orderResponse.TotalPrice = order.GrandTotalPrice
	orderResponse.TotalQuantity = order.GrandTotalQuantity
	orderResponse.PaymentStatus = order.PaymentStatus

	for _, orderDetail := range order.OrderDetails {
		orderResponse.ProductName = orderDetail.Product.Name
	}

	return orderResponse
}

func FormatterGetAllOrderUser(orders []*entities.OrderEntity) []*GetAllOrderUserResponse {
	var ordersResponse []*GetAllOrderUserResponse

	for _, order := range orders {
		orderResponse := FormatGetAllOrderUser(order)
		ordersResponse = append(ordersResponse, orderResponse)
	}

	return ordersResponse
}
