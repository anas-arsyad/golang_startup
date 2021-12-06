package transaction

type GetCampaignIdInput struct {
	CampaignId int `uri:"id" binding:"required"`
}
