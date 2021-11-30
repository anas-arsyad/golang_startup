package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(idUser int) ([]Campaign, error)
	GetCampaignById(input GetCampaignById) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(idCampaign GetCampaignById, input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetCampaigns(idUser int) ([]Campaign, error) {

	if idUser == 0 {
		campaigns, err := s.repository.FindAll()
		if err != nil {
			return campaigns, err
		}
		return campaigns, err
	}

	campaigns, err := s.repository.FindByUserId(idUser)
	if err != nil {
		return campaigns, err
	}
	return campaigns, err

}

func (s *service) GetCampaignById(input GetCampaignById) (Campaign, error) {

	campaignById, err := s.repository.FindByCampaignId(input.Id)
	if err != nil {
		return campaignById, err
	}
	return campaignById, err

}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		User:             input.User,
	}
	textComabine := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(textComabine)
	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, err

}

func (s *service) UpdateCampaign(idCampaign GetCampaignById, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByCampaignId(idCampaign.Id)
	if err != nil {
		return campaign, err
	}

	if campaign.User.ID != input.User.ID {
		return campaign, errors.New("Campaign Not Found")
	}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	updateCamapaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCamapaign, err
	}
	return updateCamapaign, err
}
