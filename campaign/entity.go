package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	CurrentAmount    int
	GoalAmount       int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignsImage   []CampaignsImage
	User             user.User
}

type CampaignsImage struct {
	ID         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
