package entities

import "time"

type OrderEntity struct {
	Id                 string               `gorm:"column:id;primaryKey" json:"id"`
	IdOrder            string               `gorm:"column:id_order;type:VARCHAR(255)" json:"id_order"`
	UserId             int                  `gorm:"column:user_id" json:"user_id"`
	GrandTotalQuantity int                  `gorm:"column:grand_total_quantity" json:"grand_total_quantity"`
	GrandTotalPrice    int                  `gorm:"column:grand_total_price" json:"grand_total_price"`
	TotalAmountPaid    int                  `gorm:"column:total_amount_paid" json:"total_amount_paid"`
	OrderStatus        string               `gorm:"column:order_status;type:VARCHAR(255)" json:"order_status"`
	PaymentStatus      string               `gorm:"column:payment_status;type:VARCHAR(255)" json:"payment_status"`
	PaymentMethod      string               `gorm:"column:payment_method;type:VARCHAR(255)" json:"payment_method"`
	CreatedAt          time.Time            `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time            `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt          *time.Time           `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	User               UserEntity           `gorm:"foreignKey:UserId" json:"user"`
	OrderDetails       []OrderDetailsEntity `gorm:"foreignKey:OrderId" json:"order_details"`
}

type OrderDetailsEntity struct {
	Id         int           `gorm:"column:id;primaryKey" json:"id"`
	OrderId    string        `gorm:"column:order_id;type:VARCHAR(255)" json:"order_id"`
	ProductId  int           `gorm:"column:product_id" json:"product_id"`
	Quantity   int           `gorm:"column:quantity" json:"quantity"`
	TotalPrice int           `gorm:"column:total_price" json:"total_price"`
	Product    ProductEntity `json:"product,omitempty" gorm:"foreignKey:ProductId"`
}

func (OrderEntity) TableName() string {
	return "orders"
}

func (OrderDetailsEntity) TableName() string {
	return "order_details"
}
