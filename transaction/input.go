package transaction

import "bwastartup/user"

type GetCampaignIdInput struct {
	CampaignId int `uri:"id" binding:"required"`
	User       user.User
}
