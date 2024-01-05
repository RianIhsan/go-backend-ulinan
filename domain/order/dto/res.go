package dto

import "ulinan/entities"

type GetAllOrderUserResponse struct {
	ID              string                 `json:"id"`
	IdOrder         string                 `json:"id_order"`
	UserID          int                    `json:"user_id"`
	GrandTotalPrice int                    `json:"grand_total_price"`
	TotalQuantity   int                    `json:"total_quantity"`
	PaymentStatus   string                 `json:"payment_status"`
	ArrivalDate     string                 `json:"arrival_date"`
	Products        []OrderProductResponse `json:"products"`
}

type OrderProductResponse struct {
	IDProduct    int    `json:"id_product"`
	ProductName  string `json:"product_name"`
	PricePerItem int    `json:"price_per_item"`
	Quantity     int    `json:"quantity"`
	TotalPrice   int    `json:"total_price"`
	ArrivalDate  string `json:"arrival_date"`
}

func FormatGetAllOrderUser(order *entities.OrderEntity) *GetAllOrderUserResponse {
	orderResponse := &GetAllOrderUserResponse{
		ID:              order.Id,
		IdOrder:         order.IdOrder,
		UserID:          order.UserId,
		GrandTotalPrice: order.GrandTotalPrice,
		TotalQuantity:   order.GrandTotalQuantity,
		PaymentStatus:   order.PaymentStatus,
		ArrivalDate:     order.ArrivalDate.Format("02 January 2006"),
	}

	var products []OrderProductResponse
	for _, orderDetail := range order.OrderDetails {
		product := OrderProductResponse{
			IDProduct:    orderDetail.Product.ID,
			ProductName:  orderDetail.Product.Name,
			PricePerItem: orderDetail.Product.Price,
			Quantity:     orderDetail.Quantity,
			TotalPrice:   orderDetail.TotalPrice,
			ArrivalDate:  order.ArrivalDate.Format("02 January 2006"),
		}
		products = append(products, product)
	}

	orderResponse.Products = products
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
