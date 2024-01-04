package payment

import (
	"errors"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateCoreAPIPaymentRequest(coreClient coreapi.Client, orderID string, totalAmountPaid int, paymentType coreapi.CoreapiPaymentType, name string, email string, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error) {
	var paymentRequest *coreapi.ChargeReq

	switch paymentType {
	case coreapi.PaymentTypeQris, coreapi.PaymentTypeGopay, coreapi.PaymentTypeBankTransfer:
		paymentRequest = &coreapi.ChargeReq{
			PaymentType: paymentType,
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderID,
				GrossAmt: int64(totalAmountPaid),
			},
		}

		if paymentType == coreapi.PaymentTypeBankTransfer {
			paymentRequest.BankTransfer = &coreapi.BankTransferDetails{
				Bank: bank,
			}
		}
	default:
		return nil, errors.New("invalid payment type")
	}

	paymentRequest.CustomerDetails = &midtrans.CustomerDetails{
		FName: name,
		Phone: phone,
		Email: email,
	}

	resp, err := coreClient.ChargeTransaction(paymentRequest)
	if err != nil {
		fmt.Println("Error creating payment request:", err.GetMessage())
		return nil, err
	}
	fmt.Println("save payment data: OrderId", orderID, "Name=", name, "Email=", email, "Phone=", phone)
	fmt.Println("Payment request created successfully:", resp)
	return resp, nil
}
