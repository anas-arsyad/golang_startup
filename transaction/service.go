package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
)

type Service interface {
	GetTransactionCampaignById(input GetCampaignIdInput) ([]Transaction, error)
	GetTransactionUserById(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(r Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{r, campaignRepository, paymentService}
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

	paymentTransaction := payment.Transaction{}
	paymentTransaction.Id = transaction.ID
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
