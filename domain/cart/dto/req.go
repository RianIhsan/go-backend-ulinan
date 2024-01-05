package dto

type AddCartItemsRequest struct {
	UserID      int    `form:"user_id" json:"user_id"`
	ProductID   int    `form:"product_id" json:"product_id" validate:"required"`
	Quantity    int    `form:"quantity" json:"quantity" validate:"required"`
	ArrivalDate string `form:"arrival_date" json:"arrival_date" validate:"required,datetime=2006-01-02"`
}

type ReduceCartItemsRequest struct {
	CartItemID int `form:"cart_item_id" json:"cart_item_id"`
	Quantity   int `form:"quantity" json:"quantity" validate:"required"`
}
