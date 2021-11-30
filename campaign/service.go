package campaign

type Service interface {
	GetCampaigns(idUser int) ([]Campaign, error)
	GetCampaignById(input GetCampaignById) (Campaign, error)
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
