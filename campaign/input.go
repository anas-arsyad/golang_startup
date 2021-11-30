package campaign

import "bwastartup/user"

type GetCampaignById struct {
	Id int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"shortDescription" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	GoalAmount       int    `json:"goalAmount" binding:"required"`
	User             user.User
}
