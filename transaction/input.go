package transaction

import "bwastartup/user"

type GetCampaignIdInput struct {
	CampaignId int `uri:"id" binding:"required"`
	User       user.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaignId" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transactionStatus"`
	OrderId           string `json:"orderId"`
	PaymentType       string `json:"paymentType"`
	FraudStatus       string `json:"fraudStatus"`
}
