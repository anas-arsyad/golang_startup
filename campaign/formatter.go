package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"userId"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	ImageUrl         string `json:"imageUrl"`
	GoalAmount       int    `json:"goalAmount"`
	CurrentAmount    int    `json:"currentUser"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageUrl:         "",
		Slug:             campaign.Slug,
	}
	if len(campaign.CampaignsImage) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignsImage[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaign []Campaign) []CampaignFormatter {
	result := []CampaignFormatter{}

	for i := 0; i < len(campaign); i++ {
		inFormat := FormatCampaign(campaign[i])
		result = append(result, inFormat)
	}
	return result
}

type CampaignFormatterDetail struct {
	ID               int                       `json:"id"`
	UserId           int                       `json:"userId"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"shortDescription"`
	Description      string                    `json:"description"`
	ImageUrl         string                    `json:"imageUrl"`
	GoalAmount       int                       `json:"goalAmount"`
	CurrentAmount    int                       `json:"currentUser"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             UserForCampaignDetail     `json:"user"`
	CampaignsImage   []ImagesForCampaignDetail `json:"images"`
}

type UserForCampaignDetail struct {
	Name   string `json:"name"`
	Avatar string `json:"avatarUrl"`
}
type ImagesForCampaignDetail struct {
	ImageUrl  string `json:"imageUrl"`
	IsPrimary bool   `json:"isPrimary"`
}

func FormatCampaignById(campaign Campaign) CampaignFormatterDetail {
	user := UserForCampaignDetail{
		Name:   campaign.User.Name,
		Avatar: campaign.User.AvatarFileName,
	}

	imagesUrl := []ImagesForCampaignDetail{}
	for _, image := range campaign.CampaignsImage {
		push := ImagesForCampaignDetail{
			ImageUrl: image.FileName,
		}
		IsPrimary := false
		if image.IsPrimary == 1 {
			IsPrimary = true
		}
		push.IsPrimary = IsPrimary

		imagesUrl = append(imagesUrl, push)
	}
	campignById := CampaignFormatterDetail{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageUrl:         "",
		Slug:             campaign.Slug,
		User:             user,
		CampaignsImage:   imagesUrl,
	}

	if len(campaign.CampaignsImage) > 0 {
		campignById.ImageUrl = campaign.CampaignsImage[0].FileName
	}

	perks := []string{}

	for _, s := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(s))
	}
	campignById.Perks = perks
	return campignById
}
