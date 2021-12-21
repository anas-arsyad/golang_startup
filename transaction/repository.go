package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetCampaignById(campaignId int) ([]Transaction, error)
	GetUserById(userId int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignById(campaignId int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("User").Where("campaign_id=?", campaignId).Order("created_at desc").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, err
}

func (r *repository) GetUserById(userId int) ([]Transaction, error) {
	transactions := []Transaction{}
	err := r.db.Preload("Campaign.CampaignsImage", "campaigns_images.is_primary=1").Where("user_id=?", userId).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, err
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, err
}
func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
