package campaign

type GetCampaignById struct {
	Id int `uri:"id" binding:"required"`
}
