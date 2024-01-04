package dto

type AddCartItemsRequest struct {
	UserID    int `form:"user_id" json:"user_id"`
	ProductID int `form:"product_id" json:"product_id" validate:"required"`
	Quantity  int `form:"quantity" json:"quantity" validate:"required"`
}
