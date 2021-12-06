package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindByCampaignId(campaignId int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignsImage) (CampaignsImage, error)
	UpdateAllImageToNonPrimary(campaignId int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Preload("CampaignsImage", "campaigns_images.is_primary=1").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, err
}
func (r *repository) FindByUserId(userId int) ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Where("user_id=?", userId).Preload("CampaignsImage", "campaigns_images.is_primary=1").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, err
}

func (r *repository) FindByCampaignId(campaignId int) (Campaign, error) {
	campaign := Campaign{}

	err := r.db.Where("id=?", campaignId).Preload("CampaignsImage").Preload("User").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, err
}

func (r *repository) Save(camapign Campaign) (Campaign, error) {
	err := r.db.Create(&camapign).Error
	if err != nil {
		return camapign, err
	}

	return camapign, nil
}
func (r *repository) Update(camapign Campaign) (Campaign, error) {
	err := r.db.Save(&camapign).Error
	if err != nil {
		return camapign, err
	}

	return camapign, nil
}

func (r *repository) CreateImage(campaignImage CampaignsImage) (CampaignsImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, err

}

func (r *repository) UpdateAllImageToNonPrimary(campaignId int) (bool, error) {
	/* UPDATE campaign_images set is_primary=false where campaign_id=1 */
	err := r.db.Model(&CampaignsImage{}).Where("campaign_id=?", campaignId).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, err
}
