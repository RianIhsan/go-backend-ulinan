package payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"ulinan/config"
)

func InitSnapMidtrans(config config.Config) coreapi.Client {
	var coreClient coreapi.Client
	coreClient.New(config.Midtrans.ServerKey, midtrans.Sandbox)
	return coreClient
}
