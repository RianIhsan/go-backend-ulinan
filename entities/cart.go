package entities

import "time"

type CartEntity struct {
	Id         int               `gorm:"column:id;primaryKey" json:"id"`
	UserId     int               `gorm:"column:user_id" json:"user_id"`
	GrandTotal int               `gorm:"column:grand_total" json:"grand_total"`
	User       *UserEntity       `gorm:"foreignKey:UserId" json:"user"`
	CartItems  []*CartItemEntity `gorm:"foreignKey:CartId" json:"cart_items,omitempty"`
}

type CartItemEntity struct {
	Id          int            `gorm:"column:id;primaryKey" json:"id"`
	CartId      int            `gorm:"column:cart_id" json:"cart_id"`
	ProductId   int            `gorm:"column:product_id" json:"product_id"`
	ArrivalDate time.Time      `gorm:"column:arrival_date;type:TIMESTAMP" json:"arrival_date,omitempty"`
	Quantity    int            `gorm:"column:quantity" json:"quantity"`
	Price       int            `gorm:"column:price" json:"price"`
	TotalPrice  int            `gorm:"column:total_price" json:"total_price"`
	Product     *ProductEntity `gorm:"foreignKey:ProductId" json:"product,omitempty"`
}

func (CartEntity) TableName() string {
	return "carts"
}

func (CartItemEntity) TableName() string {
	return "cart_items"
}
