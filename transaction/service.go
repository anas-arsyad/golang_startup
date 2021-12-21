package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type Service interface {
	GetTransactionCampaignById(input GetCampaignIdInput) ([]Transaction, error)
	GetTransactionUserById(userId int) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(r Repository, campaignRepository campaign.Repository) *service {
	return &service{r, campaignRepository}
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
