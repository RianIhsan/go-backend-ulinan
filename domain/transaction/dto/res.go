package dto

import (
	"ulinan/entities"
)

type TGetAllTransactionResponse struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	PaymentStatus string `json:"payment_status"`
	CreatedAt     string `json:"created_at"`
}

func FormatGetAllTransaction(transaction *entities.OrderEntity) *TGetAllTransactionResponse {
	transactionResponse := &TGetAllTransactionResponse{}
	transactionResponse.Id = transaction.Id
	transactionResponse.Username = transaction.User.Username
	transactionResponse.PaymentStatus = transaction.PaymentStatus
	transactionResponse.CreatedAt = transaction.CreatedAt.Format("Jan 02, 2006 03:04 pm")

	return transactionResponse
}

func FormatterGetAllTransaction(transactions []*entities.OrderEntity) []*TGetAllTransactionResponse {
	var transactionResponse []*TGetAllTransactionResponse

	for _, transaction := range transactions {
		transactionRes := FormatGetAllTransaction(transaction)
		transactionResponse = append(transactionResponse, transactionRes)
	}
	return transactionResponse
}

type TransactionDetailResponse struct {
	Id            string                  `json:"id"`
	Username      string                  `json:"username"`
	PaymentStatus string                  `json:"payment_status"`
	CreatedAt     string                  `json:"created_at"`
	GrandTotal    int                     `json:"grand_total"`
	Products      []ProductDetailResponse `json:"products"`
}

type ProductDetailResponse struct {
	Id          int    `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
}

func FormatGetTransactionById(transaction *entities.OrderEntity) *TransactionDetailResponse {
	transactionResponse := &TransactionDetailResponse{}

	transactionResponse.Id = transaction.Id
	transactionResponse.Username = transaction.User.Username
	transactionResponse.PaymentStatus = transaction.PaymentStatus
	transactionResponse.CreatedAt = transaction.CreatedAt.Format("Jan 02, 2006 03:04 pm")
	transactionResponse.GrandTotal = transaction.GrandTotalPrice
	transactionResponse.Products = make([]ProductDetailResponse, 0)

	for _, orderDetail := range transaction.OrderDetails {
		productDetail := ProductDetailResponse{
			Id:          orderDetail.Product.ID,
			ProductName: orderDetail.Product.Name,
			Price:       orderDetail.Product.Price,
			Quantity:    orderDetail.Quantity,
			TotalPrice:  orderDetail.TotalPrice,
		}

		transactionResponse.Products = append(transactionResponse.Products, productDetail)
	}

	return transactionResponse
}
