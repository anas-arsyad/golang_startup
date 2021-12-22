package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
	"strconv"

	"github.com/google/uuid"
)

type Service interface {
	GetTransactionCampaignById(input GetCampaignIdInput) ([]Transaction, error)
	GetTransactionUserById(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionCampaignById(input GetCampaignIdInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByCampaignId(input.CampaignId)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("Does't have permission to this Campaign")
	}

	transaction, err := s.repository.GetCampaignById(input.CampaignId)
	if err != nil {
		return transaction, err
	}
	return transaction, err
}

func (s *service) GetTransactionUserById(userId int) ([]Transaction, error) {
	transactions, err := s.repository.GetUserById(userId)
	if err != nil {
		return transactions, err
	}
	return transactions, err
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignId = input.CampaignId
	transaction.UserId = input.User.ID
	transaction.Status = "pending"
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}
	id := uuid.New()
	paymentTransaction := payment.Transaction{}
	paymentTransaction.Id = id.String()
	paymentTransaction.Amount = transaction.Amount

	paymentUrl, err := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}
	newTransaction.PaymentUrl = paymentUrl
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, err
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transactionId, _ := strconv.Atoi(input.OrderId)
	currentTransaction, err := s.repository.GetTransactionById(transactionId)
	if err != nil {
		return err
	}
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		currentTransaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		currentTransaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		currentTransaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(currentTransaction)
	if err != nil {
		return err
	}
	campaign, err := s.campaignRepository.FindByCampaignId(updatedTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return err
}
