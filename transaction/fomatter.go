package transaction

import (
	"time"
)

type TransactionCampaignFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

func FormatCampaignTransaction(transaction Transaction) TransactionCampaignFormatter {
	formatter := TransactionCampaignFormatter{}
	formatter.ID = transaction.CampaignId
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}
func FormatCampaignTransactions(transactions []Transaction) []TransactionCampaignFormatter {
	if len(transactions) == 0 {
		return []TransactionCampaignFormatter{}
	}

	listFormatter := []TransactionCampaignFormatter{}

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		listFormatter = append(listFormatter, formatter)
	}
	return listFormatter

}

type UserTransactionFormatter struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	Campaign  CampaignFormatter
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageUrl = ""
	if len(transaction.Campaign.CampaignsImage) > 0 {
		campaignFormatter.ImageUrl = transaction.Campaign.CampaignsImage[0].FileName
	}
	formatter.Campaign = campaignFormatter
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	listFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		listFormatter = append(listFormatter, formatter)
	}
	return listFormatter

}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.CampaignID = transaction.CampaignId
	formatter.UserID = transaction.UserId
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentUrl
	return formatter
}
