package payment

import (
	"github.com/midtrans/midtrans-go/coreapi"
	"ulinan/domain/order/dto"
)

func TransactionStatus(transactionStatusResp *coreapi.TransactionStatusResponse) dto.Status {
	var status dto.Status
	if transactionStatusResp.TransactionStatus == "capture" {
		if transactionStatusResp.FraudStatus == "challenge" {
			status.PaymentStatus = "challenge"
			status.OrderStatus = "challenge"
		} else if transactionStatusResp.FraudStatus == "accept" {
			status.PaymentStatus = "Success"
			status.OrderStatus = "Proses"
		}
	} else if transactionStatusResp.TransactionStatus == "settlement" {
		status.PaymentStatus = "Success"
		status.OrderStatus = "Proses"
	} else if transactionStatusResp.TransactionStatus == "deny" {
		status.PaymentStatus = "deny"
		status.OrderStatus = "deny"
	} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
		status.PaymentStatus = "Failed"
		status.OrderStatus = "Failed"
	} else if transactionStatusResp.TransactionStatus == "pending" {
		status.PaymentStatus = "Pending"
		status.OrderStatus = "Pending"
	}
	return status
}
