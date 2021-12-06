package transaction

type Service interface {
	GetTransactionCampaignById(input GetCampaignIdInput) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetTransactionCampaignById(input GetCampaignIdInput) ([]Transaction, error) {
	transaction, err := s.repository.GetCampaignById(input.CampaignId)
	if err != nil {
		return transaction, err
	}
	return transaction, err
}
