package dto

import (
	"ulinan/entities"
)

type CartItemFormatter struct {
	CartItemID  int    `json:"cart_item_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
	ArrivalDate string `json:"arrival_date"`
}
type CartFormatter struct {
	ID         int                 `json:"id"`
	UserID     int                 `json:"user_id"`
	GrantTotal int                 `json:"grant_total"`
	CartItems  []CartItemFormatter `json:"cart_items"`
}

type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func FormatCart(cart *entities.CartEntity) *CartFormatter {
	cartResponse := &CartFormatter{
		ID:         cart.Id,
		UserID:     cart.UserId,
		GrantTotal: cart.GrandTotal,
	}

	var cartItems []CartItemFormatter
	for _, item := range cart.CartItems {
		cartItem := CartItemFormatter{
			CartItemID:  item.Id,
			ProductName: item.Product.Name,
			Price:       item.Product.Price,
			Quantity:    item.Quantity,
			TotalPrice:  item.TotalPrice,
			ArrivalDate: item.ArrivalDate.Format("02 January 2006"),
		}
		cartItems = append(cartItems, cartItem)
	}
	cartResponse.CartItems = cartItems
	if cartResponse.CartItems == nil {
		cartResponse.GrantTotal = 0

	}
	return cartResponse
}
