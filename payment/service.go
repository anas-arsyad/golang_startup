package payment

import (
	"bwastartup/helper"
	"bwastartup/user"

	"github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	clientKey := helper.EnvVariable("CLIENT_KEY")
	serverKey := helper.EnvVariable("SERVER_KEY")
	midclient := midtrans.NewClient()
	midclient.ServerKey = serverKey
	midclient.ClientKey = clientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Id,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenRes, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenRes.RedirectURL, err
}
